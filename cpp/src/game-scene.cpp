#include "scene.h"

GameScene::GameScene(std::vector<std::shared_ptr<Player>> players) {
  this->players = players;
};

GameScene::~GameScene() {};

void GameScene::Draw() {
  for (auto player = this->players.begin(); player != this->players.end();
       player++) {
    player->get()->Draw();
  }
};
