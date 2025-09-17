#include "websocket.h"
#include <arpa/inet.h> // For inet_pton()
#include <netinet/in.h>
#include <sys/socket.h>
#include <unistd.h>

int main() {
  WebSocketClient client("ws://localhost:3000/ws");

  if (client.connect()) {
    std::cout << "Connected to WebSocket server!" << std::endl;

    // Send a message
    if (client.send_message("Hello, WebSocket!")) {
      std::cout << "Message sent successfully!" << std::endl;
    }

    // Receive response
    std::string response = client.receive_message();
    if (!response.empty()) {
      std::cout << "Received: " << response << std::endl;
    }

    client.close_connection();
  } else {
    std::cerr << "Failed to connect to WebSocket server" << std::endl;
    return 1;
  }

  return 0;
}
