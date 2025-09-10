#include "unit.h"
#include "raylib.h"

Unit::Unit(Rectangle rec) {
  rectangle = rec;
  currentObjective = nullptr;
};

Unit::~Unit() { delete currentObjective; }

void Unit::Draw(Color color) { DrawRectangleRec(rectangle, color); };

Vector2 Unit::GetPos() { return Vector2{.x = rectangle.x, .y = rectangle.y}; };
