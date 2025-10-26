#include "raylib.h"
#include "websocket.h"
#include "window.h"
#include <cstddef>
#include <memory>
#include <vector>

// TODO: Trouver pourquoi GetElements est un vecteur vide
void InputHandler(std::shared_ptr<GameScene> game_scene, Gateway *gate) {

  std::shared_ptr<Player> current_player = game_scene->GetCurrentPlayer();
  Vector2 mouse_position = GetMousePosition();
  // if (IsMouseButtonDown(MOUSE_BUTTON_RIGHT)) {
  //   std::cout << "MOUSE PRESS" << mouse_position.x << " " << mouse_position.y
  //   << std::endl; std::cout << game_scene->GetElements().size() << std::endl;
  // }

  for (auto &pair_element : game_scene->GetElements()) {
    std::shared_ptr<IScreenElement> element = pair_element.second;

    if (IsMouseButtonDown(MOUSE_BUTTON_RIGHT) && element->IsSelected()) {
      element->SetSelected(false);
      std::shared_ptr<MoveUnitOut> move_unit =
          std::make_shared<MoveUnitOut>(MoveUnitOut(
              element->GetPos(), mouse_position, "playerID", element->GetID()));
      gate->PushEvent(move_unit);
    }

    // TODO: la déselection ne fonctionne pas vraiment lorsqu'on veut
    // sélectionner un élément en mouvement
    if (CheckCollisionPointRec(mouse_position, element->GetRec()) &&
        IsMouseButtonDown(MOUSE_LEFT_BUTTON) &&
        element->GetTeam() == current_player->GetTeam()) {
      element->SetSelected(true);
    }
    // else if (IsMouseButtonDown(MOUSE_LEFT_BUTTON)) {
    //   element->SetSelected(false);
    // }
  }

  if (IsKeyPressed(KEY_P)) {
    gate->Connect();
  }

  if (IsKeyPressed(KEY_J)) {
    if (gate->GetIsConnected()) {
      std::shared_ptr<GameResquest> joinGameRequest =
          std::make_shared<GameResquest>(GameResquest("", "Game1"));
      gate->PushEvent(joinGameRequest);
    }
  }

  // if (IsKeyPressed(KEY_Q)) {
  //   std::shared_ptr<StartGameOut> start_game =
  //       std::make_shared<StartGameOut>(StartGameOut(game_scene->GetGameID()));
  //   gate->PushEvent(start_game);
  // }
}

// TODO:
// IMPORTANT POUR LES EVENT QUEUE IL VA FALLOIR FAIRE DES WRAPPERS AVEC DES
// MUTUX OU AU MOINS UN TYPE DE CONTROLE DACCES A LA MEMOIRE/VOIR SI CEST
// NECESSAIRE FAIRE MIEUX QUELE POP_AND_SEND_EVENT
void CleanEvenOutQueue(Gateway *gate) {
  if (gate->GetIsConnected()) {
    size_t queue_size = gate->GetQueueOutSize();
    while (queue_size > 0) {
      std::cout << "CLEAN OUT\n";
      gate->PopAndSendEvent();
      queue_size = gate->GetQueueOutSize();
    }
  }
}

void CleanEvenInQueue(Gateway *gate, std::shared_ptr<GameScene> game_scene) {
  if (gate->GetIsConnected()) {
    while (gate->queue_in.size() > 0) {
      // std::cout << "CLEAN IN DANS WHILE\n";
      auto event = gate->queue_in.front();
      event->HandlerEvent(game_scene);
      gate->queue_in.pop();
    }
  }
}

int main() {
  Window window = Window(900, 900, "WINDOW FROM SCENE_MANAGER");
  std::vector<std::shared_ptr<IScreenElement>> units_ptr;
  Gateway gate = Gateway();

  // Init scene
  GameScene game_scene = GameScene();

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

    InputHandler(game_shared, &gate);
    CleanEvenOutQueue(&gate);
    CleanEvenInQueue(&gate, game_shared);

    EndDrawing();
  }

  window.CloseWin();
  return 0;
}
