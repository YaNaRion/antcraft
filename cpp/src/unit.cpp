#include "unit.h"
#include "raylib.h"

Unit::Unit(Rectangle rec) {
  rectangle = rec;
  current_objective = nullptr;
  is_selected = false;
};

Unit::Unit(Vector2 pos, std::string id) {
  this->current_pos = pos;
  this->id = id;
};

Unit::~Unit() { delete current_objective; }

void Unit::Draw(Color color) {
  if (this->is_selected) {
    DrawRectangleRec(rectangle, GREEN);
  } else {
    DrawRectangleRec(rectangle, color);
  }
};

Vector2 Unit::GetPos() { return this->current_pos; };

Rectangle Unit::GetRec() { return this->rectangle; };

bool Unit::IsSelected() { return this->is_selected; };

void Unit::SetSelected(bool is_selected) { this->is_selected = is_selected; };

void Unit::SetPos(Vector2 vec2) {
  this->current_pos = vec2;
  this->rectangle.x = vec2.x * SCALE_FACTOR;
  this->rectangle.y = vec2.y * SCALE_FACTOR;
};

std::string Unit::GetID() { return this->id; }
