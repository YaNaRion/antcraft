#include <arpa/inet.h>
#include <cstring>
#include <iomanip>
#include <iostream>
#include <netinet/in.h>
#include <sstream>
#include <string>
#include <sys/socket.h>
#include <unistd.h>

class WebSocketClient {
public:
  WebSocketClient(const std::string &url);
  bool connect();
  bool send_message(const std::string &message);
  std::string receive_message();
  void close_connection();
  ~WebSocketClient();

private:
  bool perform_handshake();
  void parse_url(const std::string &url);
  int sockfd;
  std::string host;
  int port;
  std::string path;
};
