#include "scene.h"
#include <iostream>

GameScene::GameScene() {
  this->elements = {};
  this->players = {};
};

GameScene::GameScene(std::vector<std::shared_ptr<Player>> players) {
  this->players = players;
  this->elements = {};
};

GameScene::~GameScene() {};

void GameScene::Draw() {
  if (this->elements.size() > 0)
    std::cout << this->elements[0]->GetPos().x << std::endl;
  for (auto &element : elements) {
    element->Draw(RED);
  }
}

void GameScene::AddElement(std::shared_ptr<IScreenElement> elem) {
  this->elements.push_back(elem);
}
