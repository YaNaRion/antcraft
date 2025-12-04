#include "event_name.pb.h"
#include "raylib.h"
#include "scene.h"

const std::string game_backgound_texture_path =
    "./assets/game_scene_background.jpg";

GameScene::GameScene() {
  this->elements = {};
  this->players = {};
  this->backgound_texture = LoadTexture(game_backgound_texture_path.c_str());
};

GameScene::~GameScene() {};

void GameScene::Draw() {
  DrawTexture(this->backgound_texture, 0, 0, WHITE);
  DrawText("DEBUG: GAME SCENE", 10, 10, 36, WHITE);
  for (auto &element : this->elements) {
    element.second->Draw();
  }
}

void GameScene::SetCurrentPlayer(std::shared_ptr<Player> player) {
  this->current_player = player;
}

std::string GameScene::GetGameID() { return this->game_id; }

std::map<std::string, std::shared_ptr<Player>> GameScene::GetPlayers() {
  return this->players;
}

std::map<std::string, std::shared_ptr<IScreenElement>>
GameScene::GetElements() {
  return this->elements;
}

void GameScene::AddElement(std::shared_ptr<IScreenElement> elem,
                           std::string player_id) {

  this->players[player_id]->AddElement(elem);
  // this->elements.push_back(elem);
  this->elements[elem->GetElementData()->GetID()] = elem;
  // this->elements[elem->GetID()] = elem;
}

void GameScene::Update(std::vector<std::shared_ptr<IScreenElement>> vectorIN,
                       Event::GameState state) {
  this->state = state;

  // TODO: ATTENTION DE PREND PAS EN COMPTE LE CAS OU DES UNITS ONT ETE
  // ENLEVER!!!

  // TODO: FAIRE EN SORTE QU'A CHAQUE UPDATE DE SERVER, L'ELEMENT SE DEPLACE
  // PLUSIEURS FOIS DANS LE SENS DU VECTEUR DIRECTEUR FAIRE LE MEME CALCULE A
  // CHAQUE FRAME DU COTE CLIENT QUE DU COTE SERVER
  for (std::shared_ptr<IScreenElement> unit : vectorIN) {
    auto foundElement = this->elements.find(unit->GetElementData()->GetID());
    if (foundElement == this->elements.end()) {
      this->elements[unit->GetElementData()->GetID()] = unit;
    } else {
      if (unit->GetElementData()->GetCurrentObjective().get() != nullptr) {
        foundElement->second->GetElementData()->SetPos(
            unit->GetElementData()->GetPos());
        foundElement->second->GetElementData()->SetCurrentObjective(
            unit->GetElementData()->GetCurrentObjective());
        foundElement->second->GetElementData()->SetDirectionVector();
      }
    }
  }
}

std::shared_ptr<Player> GameScene::GetCurrentPlayer() {
  return this->current_player;
}
