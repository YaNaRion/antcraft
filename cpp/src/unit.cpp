#include "unit.h"
#include "raylib.h"

Unit::Unit(Vector2 pos) {
  position = pos;
  currentObjective = nullptr;
};

Unit::~Unit() { delete currentObjective; }

void Unit::Draw(Color color) {
  Rectangle rec = {
      .x = this->position.x,
      .y = this->position.y,
      .width = 10,
      .height = 10,
  };
  DrawRectangleRec(rec, color);
};

// void Unit::ChangeObjective(Vector2 vec);
//
// void Unit::Move();
