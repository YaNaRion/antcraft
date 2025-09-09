#include "player.h"
#include <memory>
#include <vector>

class IScene {
public:
  virtual void Draw() = 0;
  virtual ~IScene() {};
};

class MenuScene : public IScene {
public:
  MenuScene();
  ~MenuScene();
  void Draw() override;
};

class GameScene : public IScene {
public:
  std::vector<std::shared_ptr<Player>> players;
  GameScene(std::vector<std::shared_ptr<Player>> players);
  ~GameScene();
  void Draw() override;
};

class SceneManager {
public:
  SceneManager(std::vector<std::shared_ptr<IScene>> scenes);
  ~SceneManager();
  std::vector<std::shared_ptr<IScene>> scenes;
  std::shared_ptr<IScene> current_scene;
  void Draw();
};
