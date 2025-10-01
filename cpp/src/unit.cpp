#include "unit.h"
#include "raylib.h"

Unit::Unit(Rectangle rec) {
  this->rectangle = rec;
  this->is_selected = false;
};

Unit::Unit(Vector2 pos, std::string id) {
  this->current_pos = pos;
  this->rectangle = Rectangle{
      .x = pos.x,
      .y = pos.y,
      .width = 10,
      .height = 10,
  };
  this->id = id;
  this->is_selected = false;
};

Unit::~Unit() {}

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
  // this->rectangle.x = vec2.x * SCALE_FACTOR;
  // this->rectangle.y = vec2.y * SCALE_FACTOR;
  this->rectangle.x = vec2.x;
  this->rectangle.y = vec2.y;
};

std::string Unit::GetID() { return this->id; }
