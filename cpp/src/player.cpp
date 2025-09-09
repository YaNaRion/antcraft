#include "player.h"
#include "raylib.h"

Player::Player(Color color) { this->color = color; };

void Player::Draw() {
  for (auto unit = this->units.begin(); unit != this->units.end(); unit++) {
    unit->get()->Draw(this->color);
  }
};
