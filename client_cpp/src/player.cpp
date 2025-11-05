#include "player.h"
#include "raylib.h"

Player::Player(Color color, std::vector<std::shared_ptr<IScreenElement>> units,
               std::string player_id) {
  this->color = color;
  this->units = units;
  this->player_id = player_id;
};

Player::Player(std::string player_id, int team) {
  this->player_id = player_id;
  this->units = std::vector<std::shared_ptr<IScreenElement>>();
  this->color = COLOR_TO_TEAM_ID.find(team)->second;
  this->team = team;
};

void Player::AddElement(std::shared_ptr<IScreenElement> elem) {
  this->units.push_back(elem);
}

void Player::Draw() {
  for (auto unit = this->units.begin(); unit != this->units.end(); unit++) {
    unit->get()->Draw();
  }
};

std::vector<std::shared_ptr<IScreenElement>> Player::GetUnits() {
  return this->units;
}

int Player::GetTeam() { return this->team; }

Player::~Player() {};
