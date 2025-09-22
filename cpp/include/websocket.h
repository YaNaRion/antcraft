#include "event_name.pb.h"
#include "raylib.h"
#include <iostream>
#include <vector>
#include <websocketpp/client.hpp>
#include <websocketpp/config/asio_no_tls_client.hpp>

typedef websocketpp::client<websocketpp::config::asio_client> client;

class IEvent {
public:
  ~IEvent() {};
  virtual Event::Event CreateProtoEvent() = 0;
  virtual Vector2 GetNewPos() = 0;
};

class MoveUnitPost : public IEvent {
public:
  MoveUnitPost(Vector2 old_p, Vector2 new_p, std::string playerID,
               std::string unitID);
  ~MoveUnitPost() {};
  Event::Event CreateProtoEvent() override;
  Vector2 GetNewPos() override;

private:
  Vector2 old_pos;
  Vector2 new_pos;
  std::string playerID;
  std::string unitID;
};

class Gateway {
public:
  Gateway();
  ~Gateway();

  void Send(std::shared_ptr<IEvent> ev);
  void PushEvent(std::shared_ptr<IEvent> ev);
  void PopAndSendEvent();
  size_t GetQueueSize();
  void SendEvent(std::string eventString);
  void OnMessage(websocketpp::connection_hdl hdl, client::message_ptr msg);
  client *GetClient() { return &this->c; };

  websocketpp::connection_hdl *GetHDL() { return &this->hdl; };
  std::queue<std::shared_ptr<IEvent>> queue_in;
  std::queue<std::shared_ptr<IEvent>> queue_out;

  client c;
  websocketpp::connection_hdl hdl;

private:
};
