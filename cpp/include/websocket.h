#include "raylib.h"
#include <vector>
#include <websocketpp/client.hpp>
#include <websocketpp/config/asio_no_tls_client.hpp>

#include <iostream>

typedef websocketpp::client<websocketpp::config::asio_client> client;

class WebsocketConnection {
public:
  WebsocketConnection();
  void OnMessage(websocketpp::connection_hdl hdl, client::message_ptr msg);
  client *GetClient() { return &this->c; };
  websocketpp::connection_hdl *GetHDL() { return &this->hdl; };

private:
  client c;
  websocketpp::connection_hdl hdl;
};

class IEvent {
public:
  ~IEvent() {};
  virtual void PostEvent(WebsocketConnection *ws) = 0;
};

class MoveUnit : public IEvent {
public:
  MoveUnit(Vector2 old_p, Vector2 new_p, std::string playerID,
           std::string unitID);
  ~MoveUnit() {};
  void PostEvent(WebsocketConnection *ws) override;

private:
  Vector2 old_pos;
  Vector2 new_pos;
  std::string playerID;
  std::string unitID;
};

class Gateway {
public:
  Gateway() : ws(WebsocketConnection()) {};
  ~Gateway();

  void Send(std::shared_ptr<IEvent> ev);

  void PushEvent(std::shared_ptr<IEvent> ev);
  void PopEvent();
  size_t GetQueueSize();
  std::queue<std::shared_ptr<IEvent>> queue;

  // A mettre dans private avec get
  WebsocketConnection ws;

private:
};
