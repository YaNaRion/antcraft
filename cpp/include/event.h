#include "scene.h"
#include <string>

class IEventIN {
public:
  ~IEventIN() {};
  virtual void HandlerEvent(std::shared_ptr<GameScene> game_scene) = 0;
};

class UpdateMapStateEvent : public IEventIN {
public:
  UpdateMapStateEvent(std::vector<std::shared_ptr<IScreenElement>> elements,
                      Event::GameState state);
  ~UpdateMapStateEvent() {};
  void HandlerEvent(std::shared_ptr<GameScene> gameScene) override;

private:
  std::vector<std::shared_ptr<IScreenElement>> elements;
  Event::GameState state;
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
