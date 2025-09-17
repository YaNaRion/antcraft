#include "websocket.h"
#include <arpa/inet.h>
#include <netinet/in.h>
#include <openssl/bio.h>
#include <openssl/buffer.h>
#include <openssl/evp.h>
#include <openssl/sha.h>
#include <random>
#include <string>
#include <sys/socket.h>
#include <unistd.h>

std::string base64_encode(const std::string &input) {
  BIO *bio, *b64;
  BUF_MEM *bufferPtr;

  b64 = BIO_new(BIO_f_base64());
  bio = BIO_new(BIO_s_mem());
  bio = BIO_push(b64, bio);

  BIO_set_flags(bio, BIO_FLAGS_BASE64_NO_NL);
  BIO_write(bio, input.c_str(), input.length());
  BIO_flush(bio);
  BIO_get_mem_ptr(bio, &bufferPtr);

  std::string result(bufferPtr->data, bufferPtr->length);

  BIO_free_all(bio);
  return result;
}

std::string sha1(const std::string &input) {
  unsigned char hash[SHA_DIGEST_LENGTH];
  SHA1(reinterpret_cast<const unsigned char *>(input.c_str()), input.length(),
       hash);

  return std::string(reinterpret_cast<char *>(hash), SHA_DIGEST_LENGTH);
}

std::string generate_websocket_key() {
  std::random_device rd;
  std::mt19937 gen(rd());
  std::uniform_int_distribution<> dis(0, 255);

  std::string key;
  for (int i = 0; i < 16; ++i) {
    key += static_cast<char>(dis(gen));
  }

  return base64_encode(key);
}
void WebSocketClient::parse_url(const std::string &url) {
  // Default values
  host = "localhost";
  port = 80;
  path = "/";

  // Check if it's a WebSocket URL
  if (url.find("ws://") == 0) {
    std::string remaining = url.substr(5); // Remove "ws://"

    // Find host:port part
    size_t slash_pos = remaining.find('/');
    std::string host_port;

    if (slash_pos != std::string::npos) {
      host_port = remaining.substr(0, slash_pos);
      path = remaining.substr(slash_pos);
    } else {
      host_port = remaining;
    }

    // Parse host and port
    size_t colon_pos = host_port.find(':');
    if (colon_pos != std::string::npos) {
      host = host_port.substr(0, colon_pos);
      port = std::stoi(host_port.substr(colon_pos + 1));
    } else {
      host = host_port;
    }
  } else {
    // Assume it's already just host:port/path
    std::cout << "Warning: URL doesn't start with ws://, using as host"
              << std::endl;
    host = url;
  }
}

bool WebSocketClient::connect() {
  // Create socket
  sockfd = socket(AF_INET, SOCK_STREAM, 0);
  if (sockfd < 0) {
    std::cerr << "Socket creation failed" << std::endl;
    return false;
  }

  // Set up server address
  sockaddr_in server_addr{};
  server_addr.sin_family = AF_INET;
  server_addr.sin_port = htons(port);

  // Use localhost (127.0.0.1)
  server_addr.sin_addr.s_addr = inet_addr("127.0.0.1");

  if (server_addr.sin_addr.s_addr == INADDR_NONE) {
    std::cerr << "Invalid address" << std::endl;
    close(sockfd);
    return false;
  }

  std::cout << "Connecting to " << host << ":" << port << " at path " << path
            << std::endl;

  // Connect to server
  if (::connect(sockfd, (sockaddr *)&server_addr, sizeof(server_addr)) < 0) {
    std::cerr << "Connection failed to localhost:" << port << std::endl;
    std::cerr << "Make sure a WebSocket server is running on that port"
              << std::endl;
    close(sockfd);
    return false;
  }

  return perform_handshake();
}

bool WebSocketClient::perform_handshake() {
  std::string key = generate_websocket_key();

  // Build handshake request
  std::string request = "GET " + path +
                        " HTTP/1.1\r\n"
                        "Host: " +
                        host + ":" + std::to_string(port) +
                        "\r\n"
                        "Upgrade: websocket\r\n"
                        "Connection: Upgrade\r\n"
                        "Sec-WebSocket-Key: " +
                        key +
                        "\r\n"
                        "Sec-WebSocket-Version: 13\r\n"
                        "\r\n";

  std::cout << "Sending handshake request..." << std::endl;

  // Send handshake request
  if (send(sockfd, request.c_str(), request.length(), 0) !=
      (ssize_t)request.length()) {
    std::cerr << "Handshake request failed" << std::endl;
    return false;
  }

  // Receive and validate handshake response
  char buffer[1024];
  ssize_t bytes_received = recv(sockfd, buffer, sizeof(buffer) - 1, 0);
  if (bytes_received <= 0) {
    std::cerr << "Handshake response failed" << std::endl;
    return false;
  }

  buffer[bytes_received] = '\0';
  std::string response(buffer);

  std::cout << "Received handshake response:" << std::endl;
  std::cout << response << std::endl;

  // Check if handshake was successful
  if (response.find("HTTP/1.1 101") == std::string::npos) {
    std::cerr << "Invalid handshake response - expected HTTP 101" << std::endl;
    return false;
  }

  std::cout
      << "WebSocket connection established successfully to ws://localhost:"
      << port << path << std::endl;
  return true;
}

bool WebSocketClient::send_message(const std::string &message) {
  // Simple WebSocket frame (text message, no masking for client→server)
  std::vector<unsigned char> frame;

  // FIN bit set, opcode 1 (text frame)
  frame.push_back(0x81);

  // Payload length
  if (message.length() <= 125) {
    frame.push_back(message.length());
  } else if (message.length() <= 65535) {
    frame.push_back(126);
    frame.push_back((message.length() >> 8) & 0xFF);
    frame.push_back(message.length() & 0xFF);
  } else {
    // For very long messages (not implemented here)
    return false;
  }

  // Add payload
  for (char c : message) {
    frame.push_back(c);
  }

  return send(sockfd, frame.data(), frame.size(), 0) == (ssize_t)frame.size();
}

std::string WebSocketClient::receive_message() {
  char header[2];
  ssize_t bytes_received = recv(sockfd, header, 2, 0);

  if (bytes_received != 2) {
    return "";
  }

  unsigned char opcode = header[0] & 0x0F;
  bool masked = (header[1] & 0x80) != 0;
  uint64_t payload_length = header[1] & 0x7F;

  // Handle extended payload lengths
  if (payload_length == 126) {
    char extended_len[2];
    recv(sockfd, extended_len, 2, 0);
    payload_length = (extended_len[0] << 8) | extended_len[1];
  } else if (payload_length == 127) {
    char extended_len[8];
    recv(sockfd, extended_len, 8, 0);
    // For very long messages (simplified)
    payload_length = 0;
    for (int i = 0; i < 8; i++) {
      payload_length = (payload_length << 8) | extended_len[i];
    }
  }

  // Read masking key if present
  char masking_key[4];
  if (masked) {
    recv(sockfd, masking_key, 4, 0);
  }

  // Read payload
  std::string payload;
  payload.resize(payload_length);
  recv(sockfd, &payload[0], payload_length, 0);

  // Unmask payload if necessary
  if (masked) {
    for (size_t i = 0; i < payload_length; i++) {
      payload[i] ^= masking_key[i % 4];
    }
  }

  return payload;
}

void WebSocketClient::close_connection() {
  if (sockfd != -1) {
    close(sockfd);
    sockfd = -1;
  }
}

WebSocketClient::~WebSocketClient() { close_connection(); }

WebSocketClient::WebSocketClient(const std::string &url) { parse_url(url); };
