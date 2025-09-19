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

class Event {};

class Gateway {
public:
  Gateway();
  ~Gateway();

  void PushEvent(Event ev);
  void PopEvent();
  std::queue<std::pair<Event, void *>> queue;

private:
  std::shared_ptr<WebsocketConnection> ws;
};
