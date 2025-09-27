#include "event_name.pb.h"
#include "scene.h"
#include <iostream>

GameScene::GameScene() {
  this->elements = {};
  this->players = {};
};

GameScene::~GameScene() {};

void GameScene::Draw() {
  for (auto &player : players) {
    player.second->Draw();
  }
}

std::map<std::string, std::shared_ptr<IScreenElement>>
GameScene::GetElements() {
  return this->elements;
}

void GameScene::AddElement(std::shared_ptr<IScreenElement> elem,
                           std::string player_id) {
  this->players[player_id]->AddElement(elem);
  this->elements[elem->GetID()] = elem;
}

void GameScene::Update(std::vector<std::shared_ptr<IScreenElement>> vectorIN,
                       Event::GameState state) {
  this->state = state;

  // TODO: ATTENTION DE PREND PAS EN COMPTE LE CAS OU DES UNITS ONT ETE
  // ENLEVER!!!
  for (auto &unit : vectorIN) {
    auto foundElement = this->elements.find(unit->GetID());
    if (foundElement == this->elements.end()) {
      this->elements[unit->GetID()] = unit;
    } else {
      foundElement->second->SetPos(unit->GetPos());
    }
  }
}
