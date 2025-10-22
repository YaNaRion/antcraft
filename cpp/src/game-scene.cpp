#include "event_name.pb.h"
#include "scene.h"

GameScene::GameScene() {
  this->elements = {};
  this->players = {};
};

GameScene::~GameScene() {};

// TODO : AJOUER PARTOUT LA TEAM DE UNIT DANS SES DONNÉES A LUI POUR SIMPLIFIER
// LA CHOSE
void GameScene::Draw() {
  // for (auto &element : this->elements) {
  //   element.second->Draw(this->current_player->GetTeamColor());
  // }

  for (auto &player : this->players) {
    for (auto &element : this->elements) {
      element.second->Draw(player.second->GetTeamColor());
    }
  }
}

void GameScene::SetCurrentPlayer(std::shared_ptr<Player> player) {
  this->current_player = player;
}

void GameScene::SetGameID(std::string gameID) { this->game_id = gameID; }

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
  this->elements[elem->GetID()] = elem;
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
    auto foundElement = this->elements.find(unit->GetID());
    if (foundElement == this->elements.end()) {
      this->elements[unit->GetID()] = unit;
    } else {
      if (unit->GetCurrentObjective().get() != nullptr) {
        foundElement->second->SetCurrentObjective(unit->GetCurrentObjective());
        foundElement->second->SetDirectionVector();
      }
    }
  }
}
