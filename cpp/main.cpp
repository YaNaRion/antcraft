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
      std::shared_ptr<MoveUnitPost> move_unit =
          std::make_shared<MoveUnitPost>(MoveUnitPost(
              element->GetPos(), mouse_position, "playerID", "unitID"));

      gate->PushEvent(move_unit);
    }

    if (CheckCollisionPointRec(mouse_position, element->GetRec()) &&
        IsMouseButtonDown(MOUSE_LEFT_BUTTON)) {
      element->SetSelected(true);
    }
  }
}

void threadGate() { std::cout << "DANS THREAD\n"; }

// IMPORTANT POUR LES EVENT QUEUE IL VA FALLOIR FAIRE DES WRAPPERS AVEC DES
// MUTUX OU AU MOINS UN TYPE DE CONTROLE DACCES A LA MEMOIRE/VOIR SI CEST
// NECESSAIRE
void CleanEvenOutQueue(Gateway *gate) {
  size_t queue_size = gate->GetQueueSize();
  while (queue_size > 0) {
    std::cout << "CLEAN OUT\n";
    gate->PopEvent();
    queue_size = gate->GetQueueSize();
  }
}

void CleanEvenInQueue(Gateway *gate, GameScene *game_scene) {
  while (gate->queue_in.size() > 0) {
    std::cout << "CLEAN IN\n";
    auto event = gate->queue_in.front();
    std::shared_ptr<IScreenElement> unit =
        game_scene->players.front()->GetUnit().front();
    unit->SetPos(event->GetNewPos());
    std::cout << "NOUVELLE POSITION RECU DU SERVER ET RECUP DANS CLEAN EVENT "
                 "IN QUEUE \t X: "
              << unit->GetPos().x << "Y: " << unit->GetPos().y << std::endl;
    gate->queue_in.pop();
  }
}

int main() {
  Window window = Window(1920, 1080, "WINDOW FROM SCENE_MANAGER");

  Gateway *gate = new Gateway();

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

  std::cout << "Initialisation done" << std::endl;
  while (!window.ShouldWindowClose()) {
    BeginDrawing();
    ClearBackground(BLACK);
    scene_manager.Draw();

    InputHandler(units_ptr, gate);
    CleanEvenOutQueue(gate);
    CleanEvenInQueue(gate, &game_scene);

    EndDrawing();
  }

  window.CloseWin();

  return 0;
}
