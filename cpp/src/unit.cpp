#include "unit.h"
#include "raylib.h"
#include <cmath>
// #include <iostream>
// #include <print>

Unit::Unit(Rectangle rec) {
  this->rectangle = rec;
  this->is_selected = false;
};

Unit::Unit(Vector2 pos, Vector2 currentObj, std::string id, int team) {
  this->current_objective = std::make_shared<Vector2>(currentObj);
  this->current_pos = pos;
  this->rectangle = Rectangle{
      .x = pos.x,
      .y = pos.y,
      .width = 10,
      .height = 10,
  };
  this->id = id;
  this->is_selected = false;
  this->team = team;
  this->color = COLOR_TO_TEAM_ID.find(team)->second;
};

Unit::~Unit() {}

void Unit::Draw() {
  this->MoveUnitPerFrame();
  if (this->is_selected) {
    DrawRectangleRec(rectangle, GREEN);
  } else {
    DrawRectangleRec(rectangle, this->color);
  }
};

Vector2 Unit::GetPos() { return this->current_pos; };

Rectangle Unit::GetRec() { return this->rectangle; };

bool Unit::IsSelected() { return this->is_selected; };

int Unit::GetTeam() { return this->team; };

void Unit::SetSelected(bool is_selected) { this->is_selected = is_selected; };

void Unit::SetPos(Vector2 vec2) {
  this->current_pos = vec2;
  // this->rectangle.x = vec2.x * SCALE_FACTOR;
  // this->rectangle.y = vec2.y * SCALE_FACTOR;
  this->rectangle.x = vec2.x;
  this->rectangle.y = vec2.y;
};

void Unit::SetCurrentObjective(Vector2 vec2) {
  this->current_objective = std::make_shared<Vector2>(vec2);
};

void Unit::SetCurrentObjective(std::shared_ptr<Vector2> vec2) {
  this->current_objective = vec2;
};

std::string Unit::GetID() { return this->id; }

void Unit::MoveUnitPerFrame() {
  if (this->direction_vector_normalize.get() != nullptr) {
    // std::printf("CURRENT POS: X = %f, Y = %f \n", this->current_pos.x,
    //             this->current_pos.y);
    // std::printf("CURRENT OBJ: X = %f, Y = %f \n", this->current_objective->x,
    //             this->current_objective->y);
    if ((this->current_pos.x == this->current_objective->x &&
         this->current_pos.y == this->current_objective->y) ||
        (this->current_objective->x < 0 && this->current_objective->y < 0)) {
      return;
    }
    const float speed = 1.0;
    float moveX = this->direction_vector_normalize->x * speed;
    float moveY = this->direction_vector_normalize->y * speed;

    Vector2 toTarget = Vector2{
        .x = current_objective->x - this->current_pos.x,
        .y = current_objective->y - this->current_pos.y,
    };

    // std::cout << "MOVEMENT DES PIECES AVANT IF COLLIDE\n";
    // std::cout << this->current_pos.x << std::endl;
    // std::cout << this->current_pos.y << std::endl;
    if ((moveX * toTarget.x + moveY * toTarget.y) <= 0) {
      moveX = toTarget.x;
      moveY = toTarget.y;
    }

    this->current_pos.x += (float)moveX;
    this->current_pos.y += (float)moveY;

    this->rectangle.x = this->current_pos.x;
    this->rectangle.y = this->current_pos.y;

    // std::cout << "MOVEMENT DES PIECES\n";
    // std::cout << this->current_pos.x << std::endl;
    // std::cout << this->current_pos.y << std::endl;
    // std::printf("VALEUR DE CURRENT OBJ: x = %f, y = %f\n",
    //             this->GetCurrentObjective()->x,
    //             this->GetCurrentObjective()->y);
    // std::printf("VALEUR DE MOUVEMENT: x += %d, y += %d\n", moveX, moveY);
  };
}

std::shared_ptr<Vector2> Unit::GetCurrentObjective() {
  return this->current_objective;
}

void Unit::SetDirectionVector() {
  Vector2 direction_vector =
      Vector2{.x = this->current_objective->x - this->current_pos.x,
              .y = this->current_objective->y - this->current_pos.y};
  float lenght =
      pow((pow(direction_vector.x, 2) + pow(direction_vector.y, 2)), 0.5);

  direction_vector.x = direction_vector.x / lenght;
  direction_vector.y = direction_vector.y / lenght;
  this->direction_vector_normalize =
      std::make_shared<Vector2>(direction_vector);
};
