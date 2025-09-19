#include "raylib.h"
#include "websocket.h"
#include "window.h"
#include <memory>
#include <vector>

void InputHandler(std::vector<std::shared_ptr<IScreenElement>> element_ptr,
                  Gateway *gate) {
  for (std::shared_ptr<IScreenElement> element : element_ptr) {
    Vector2 mouse_position = GetMousePosition();
    if (IsMouseButtonDown(MOUSE_BUTTON_RIGHT) && element->IsSelected()) {
      element->SetSelected(false);
      element->SetPos(mouse_position);
      std::shared_ptr<MoveUnit> move_unit = std::make_shared<MoveUnit>(
          MoveUnit(element->GetPos(), mouse_position, "playerID", "unitID"));

      gate->PushEvent(move_unit);
    }

    if (CheckCollisionPointRec(mouse_position, element->GetRec()) &&
        IsMouseButtonDown(MOUSE_LEFT_BUTTON)) {
      element->SetSelected(true);
    }
  }
}

void threadGate() { std::cout << "DANS THREAD\n"; }

void CleanEvenQueue(Gateway *gate) {
  while (gate->PopEvent() != 0) {
    std::cout << "DANS CLEAN";
  }
}

int main() {
  Window window = Window(1920, 1080, "WINDOW FROM SCENE_MANAGER");

  Gateway *gate = new Gateway();
  std::cout << "Initialisation done" << std::endl;

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

    if (gate != nullptr) {
      InputHandler(units_ptr, gate);
      CleanEvenQueue(gate);
    }

    scene_manager.Draw();
    EndDrawing();
  }

  window.CloseWin();

  return 0;
}
