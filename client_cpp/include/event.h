#include "scene.h"
#include <string>

class IEventIN {
public:
  ~IEventIN() {};
  virtual void HandlerEvent(std::shared_ptr<GameScene> game_scene) = 0;
};

class UpdateMapStateEventIN : public IEventIN {
public:
  UpdateMapStateEventIN(std::vector<std::shared_ptr<IScreenElement>> elements,
                        Event::GameState state);
  ~UpdateMapStateEventIN() {};
  void HandlerEvent(std::shared_ptr<GameScene> gameScene) override;

private:
  std::vector<std::shared_ptr<IScreenElement>> elements;
  Event::GameState state;
};

class JoinGameEventIN : public IEventIN {
public:
  JoinGameEventIN(std::string game_id, std::string player_id, int team);
  ~JoinGameEventIN() {};
  void HandlerEvent(std::shared_ptr<GameScene> gameScene) override;

private:
  std::string game_id;
  std::string player_id;
  int team;
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

class StartGameOut : public IEventOut {
public:
  StartGameOut(std::string gameID);
  ~StartGameOut() {};
  Event::Event CreateProtoEvent() override;

private:
  std::string game_id;
};
