#include "raylib.h"
#include "websocket.h"
#include "window.h"
#include <cstddef>
#include <memory>
#include <string>
#include <vector>

bool isUnitSelect = false;
std::shared_ptr<IScreenElement> selected_element;

bool checkCollide(std::map<std::string, std::shared_ptr<IScreenElement>> vec,
                  std::shared_ptr<IScreenElement> attacked_elements,
                  Vector2 mouse_pos) {
  for (auto elem : vec) {
    if (elem.second->GetElementData()->GetID() ==
        attacked_elements->GetElementData()->GetID()) {
      continue;
    }
    if (CheckCollisionPointRec(mouse_pos,
                               elem.second->GetElementData()->GetRec())) {
      return true;
    }
  }
  return false;
}

// TODO: Trouver pourquoi GetElements est un vecteur vide
void InputHandler(SceneManager *scene_manager, Gateway *gate) {
  std::shared_ptr<Player> current_player =
      scene_manager->game_scene->GetCurrentPlayer();
  Vector2 mouse_position = GetMousePosition();

  for (auto &pair_element : scene_manager->game_scene->GetElements()) {
    std::shared_ptr<IScreenElement> element = pair_element.second;

    if (IsMouseButtonDown(MOUSE_BUTTON_RIGHT) && isUnitSelect &&
        checkCollide(scene_manager->game_scene->GetElements(), element,
                     mouse_position) &&
        element->GetElementData()->IsSelected()) {
      std::shared_ptr<AttackUnit> attack_unit =
          std::make_shared<AttackUnit>(AttackUnit(
              scene_manager->game_scene->GetGameID(), selected_element, element,
              scene_manager->game_scene->GetCurrentPlayer()->GetID(),
              scene_manager->game_scene->GetCurrentPlayer()->GetTeam()));
      gate->PushEvent(attack_unit);
    }

    if (IsMouseButtonDown(MOUSE_BUTTON_RIGHT) &&
        element->GetElementData()->IsSelected()) {
      element->GetElementData()->SetSelected(false);
      isUnitSelect = false;
      std::shared_ptr<MoveUnitOut> move_unit = std::make_shared<MoveUnitOut>(
          MoveUnitOut(element->GetElementData()->GetPos(), mouse_position,
                      "playerID", element->GetElementData()->GetID()));
      gate->PushEvent(move_unit);
    }

    // TODO: la déselection ne fonctionne pas vraiment lorsqu'on veut
    // sélectionner un élément en mouvement
    if (CheckCollisionPointRec(mouse_position,
                               element->GetElementData()->GetRec()) &&
        IsMouseButtonDown(MOUSE_LEFT_BUTTON) &&
        element->GetElementData()->GetTeam() == current_player->GetTeam()) {
      element->GetElementData()->SetSelected(true);
      isUnitSelect = true;
      selected_element = element;
    }
    // else if (IsMouseButtonDown(MOUSE_LEFT_BUTTON)) {
    //   element->SetSelected(false);
    // }
  }

  if (IsKeyPressed(KEY_P) ||
      (CheckCollisionPointRec(
           GetMousePosition(),
           scene_manager->menu_scene->play_button.GetRect()) &&
       IsMouseButtonReleased(MOUSE_BUTTON_LEFT))) {
    scene_manager->current_scene = scene_manager->game_scene;
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
  Window window = Window(1920, 1080, "WINDOW FROM SCENE_MANAGER");
  InitAudioDevice();

  std::vector<std::shared_ptr<IScreenElement>> units_ptr;
  Gateway gate = Gateway();
  // Init scene
  GameScene game_scene = GameScene();

  // Create MenuScene
  MenuScene menu;
  menu.SetBackgoundTexture();

  std::shared_ptr<MenuScene> menu_shared = std::make_shared<MenuScene>(menu);
  std::shared_ptr<GameScene> game_shared =
      std::make_shared<GameScene>(game_scene);

  std::vector<std::shared_ptr<IScene>> vectorScene;
  vectorScene.push_back(menu_shared);
  vectorScene.push_back(game_shared);
  SceneManager scene_manager = SceneManager(game_shared, menu_shared);

  Music menu_theme = LoadMusicStream("./assets/sound/Foundation Main "
                                     "Theme [Official Music Video].mp3");

  PlayMusicStream(menu_theme);

  std::cout << "Initialisation done" << std::endl;
  while (!window.ShouldWindowClose()) {
    Vector2 mouse_position = GetMousePosition();
    BeginDrawing();
    ClearBackground(WHITE);
    scene_manager.Draw();

    UpdateMusicStream(menu_theme);

    std::string mouse_pos_string =
        "MOUSE POSITION X: " + std::to_string(mouse_position.x) +
        "MOUSE POSITION Y: " + std::to_string(mouse_position.y);
    DrawText(mouse_pos_string.c_str(), mouse_position.x, mouse_position.y, 12,
             WHITE);

    InputHandler(&scene_manager, &gate);
    CleanEvenOutQueue(&gate);
    CleanEvenInQueue(&gate, game_shared);

    EndDrawing();
  }

  Vector2 mouse_position = GetMousePosition();
  std::string text = "Position en X" + std::to_string(mouse_position.x) +
                     "Postion en Y" + std::to_string(mouse_position.y);
  DrawText(text.c_str(), 10, 10, 12, WHITE);

  window.CloseWin();
  return 0;
}
