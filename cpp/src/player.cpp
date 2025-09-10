#include "player.h"
#include "raylib.h"

Player::Player(Color color, std::vector<std::shared_ptr<IUnit>> units) {
  this->color = color;
  this->units = units;
};

void Player::Draw() {
  for (auto unit = this->units.begin(); unit != this->units.end(); unit++) {
    unit->get()->Draw(this->color);
  }
};

Player::~Player() {};
