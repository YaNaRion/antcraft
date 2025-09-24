#include "event_name.pb.h"
#include "raylib.h"
#include "scene.h"
#include <iostream>
#include <vector>
#include <websocketpp/client.hpp>
#include <websocketpp/config/asio_no_tls_client.hpp>

typedef websocketpp::client<websocketpp::config::asio_client> client;

class IEventIN {
public:
  ~IEventIN() {};
  virtual void HandlerEvent(std::shared_ptr<GameScene> game_scene) = 0;
};

class AddUnitIN : public IEventIN {
public:
  AddUnitIN(Rectangle rec);
  ~AddUnitIN() {};
  void HandlerEvent(std::shared_ptr<GameScene> gameScene) override;

private:
  Rectangle rec;
};

class MoveUnitIN : public IEventIN {
public:
  MoveUnitIN(Vector2 old_p, Vector2 new_p, std::string playerID,
             std::string unitID);
  ~MoveUnitIN() {};
  void HandlerEvent(std::shared_ptr<GameScene> gameScene) override;

private:
  Vector2 old_pos;
  Vector2 new_pos;
  std::string playerID;
  std::string unitID;
};

class IEventOut {
public:
  ~IEventOut() {};
  virtual Event::Event CreateProtoEvent() = 0;
};

class GameResquest : public IEventOut {
public:
  GameResquest(std::string playerID, std::string gameID);
  ~GameResquest() {};
  Event::Event CreateProtoEvent() override;

private:
  std::string playerID;
  std::string gameID;
};

class MoveUnitOut : public IEventOut {
public:
  MoveUnitOut(Vector2 old_p, Vector2 new_p, std::string playerID,
              std::string unitID);
  ~MoveUnitOut() {};
  Event::Event CreateProtoEvent() override;
  Vector2 GetNewPos();

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

  void Send(std::shared_ptr<IEventOut> ev);
  void PushEvent(std::shared_ptr<IEventOut> ev);
  void PopAndSendEvent();
  size_t GetQueueOutSize();
  void SendEvent(std::string eventString);
  void OnMessage(websocketpp::connection_hdl hdl, client::message_ptr msg);
  client *GetClient() { return &this->c; };

  websocketpp::connection_hdl *GetHDL() { return &this->hdl; };
  std::queue<std::shared_ptr<IEventIN>> queue_in;
  std::queue<std::shared_ptr<IEventOut>> queue_out;

  client c;
  websocketpp::connection_hdl hdl;

private:
};
