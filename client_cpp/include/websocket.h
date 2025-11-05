#include "event.h"
#include "event_name.pb.h"
#include "raylib.h"
#include <iostream>
#include <vector>
#include <websocketpp/client.hpp>
#include <websocketpp/config/asio_no_tls_client.hpp>

typedef websocketpp::client<websocketpp::config::asio_client> client;

class Gateway {
public:
  Gateway();
  ~Gateway();

  void Send(std::shared_ptr<IEventOut> ev);
  void PushEvent(std::shared_ptr<IEventOut> ev);
  void PopAndSendEvent();
  size_t GetQueueOutSize();
  void SendEvent(std::string eventString);
  void OnMessage(websocketpp::connection_hdl hdl, client::message_ptr msg);
  client *GetClient() { return &this->c; };
  void Connect();
  bool GetIsConnected();
  void SyncGameEventHandler(const Event::SyncGameState &game_state);
  websocketpp::connection_hdl *GetHDL() { return &this->hdl; };

  std::queue<std::shared_ptr<IEventIN>> queue_in;
  std::queue<std::shared_ptr<IEventOut>> queue_out;

private:
  client c;
  websocketpp::connection_hdl hdl;
  bool IsConnected;
};
