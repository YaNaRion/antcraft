#include <vector>
#include <websocketpp/client.hpp>
#include <websocketpp/config/asio_no_tls_client.hpp>

#include <iostream>

typedef websocketpp::client<websocketpp::config::asio_client> client;

class WebsocketConnection {
public:
  WebsocketConnection();
  void OnMessage(websocketpp::connection_hdl hdl, client::message_ptr msg);

private:
  client c;
};
