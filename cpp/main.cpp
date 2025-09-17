#include "raylib.h"
#include "window.h"
#include <memory>
#include <queue>
#include <vector>

struct Event {};

struct InputQueue {
  void PushEvent();
  void PopEvent();

  std::queue<Event> queue;
};

void InputHandler() {}

int main() {

  Window window = Window(1920, 1080, "WINDOW FROM SCENE_MANAGER");

  Rectangle rec = {
      .x = 400,
      .y = 400,
      .width = 10,
      .height = 10,
  };
  Unit unit = Unit(rec); // Create 1st player

  std::shared_ptr<Unit> posPtr = std::make_shared<Unit>(unit);
  std::vector<std::shared_ptr<IScreenElement>> units_ptr;
  units_ptr.push_back(posPtr);

  Player player = Player(RED, units_ptr);

  std::shared_ptr<Player> player_ptr = std::make_shared<Player>(player);
  std::vector<std::shared_ptr<Player>> players;
  players.push_back(player_ptr);

  // Init scene
  GameScene game_scene = GameScene(players);
  // Create MenuScene
  MenuScene menu = MenuScene();

  std::shared_ptr<MenuScene> menu_shared = std::make_shared<MenuScene>(menu);
  std::shared_ptr<GameScene> game_shared =
      std::make_shared<GameScene>(game_scene);

  std::vector<std::shared_ptr<IScene>> vectorScene;
  vectorScene.push_back(menu_shared);
  vectorScene.push_back(game_shared);
  SceneManager scene_manager = SceneManager(vectorScene);

  while (!window.ShouldWindowClose()) {
    BeginDrawing();
    ClearBackground(BLACK);

    scene_manager.Draw();
    EndDrawing();
  }

  window.CloseWin();

  return 0;
}
