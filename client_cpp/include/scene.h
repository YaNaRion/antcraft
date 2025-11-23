#include "event_name.pb.h"
#include "player.h"
#include <map>
#include <memory>
#include <utility>
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
  GameScene();
  // GameScene(std::vector<std::shared_ptr<Player>> players);
  ~GameScene();
  void AddElement(std::shared_ptr<IScreenElement>, std::string player_id);
  void PlayerElement(std::shared_ptr<Player>);
  void GetElementByID(std::string elementID);
  void SetCurrentPlayer(std::shared_ptr<Player> player);
  std::shared_ptr<Player> GetCurrentPlayer();
  std::string GetGameID();
  void SetGameID(std::string gameID);
  void Draw() override;
  void Update(std::vector<std::shared_ptr<IScreenElement>> vectorIN,
              Event::GameState state);
  std::map<std::string, std::shared_ptr<Player>> GetPlayers();

  std::map<std::string, std::shared_ptr<IScreenElement>> GetElements();

private:
  // Key is ElementID, UnitID
  std::map<std::string, std::shared_ptr<IScreenElement>> elements;

  std::map<std::pair<int, int>, std::shared_ptr<IScreenElement>> elements_pair;
  // std::vector<std::shared_ptr<IScreenElement>> elements;
  std::map<std::string, std::shared_ptr<Player>> players;
  Event::GameState state;
  std::shared_ptr<Player> current_player;
  std::string game_id;
};

class SceneManager {
public:
  SceneManager(std::vector<std::shared_ptr<IScene>> scenes);
  ~SceneManager();
  std::vector<std::shared_ptr<IScene>> scenes;
  std::shared_ptr<IScene> current_scene;
  void Draw();
};
