#include "unit.h"
#include "raylib.h"
#include <cmath>
#include <iostream>
// #include <print>

Building::Building(Vector2 pos, Vector2 currentObj, Vector2 size,
                   std::string id, int team) {
  this->data = std::make_shared<ElementData>(
      ElementData(pos, currentObj, size, id, team));
}

void Building::Draw() {
  std::cout << "UNIT POS X: " << this->data->GetPos().x;
  std::cout << "UNIT POS Y: " << this->data->GetPos().y << std::endl;

  if (this->data->IsSelected()) {
    DrawRectangleRec(this->data->GetRec(), GREEN);
  } else {
    DrawRectangleRec(this->data->GetRec(), this->data->color);
  }
};

ScreenElementType Building::GetType() { return ScreenElementType::Base; }
std::shared_ptr<ElementData> Building::GetElementData() { return this->data; };

Unit::Unit(Vector2 pos, Vector2 currentObj, Vector2 size, std::string id,
           int team) {
  this->data = std::make_shared<ElementData>(
      ElementData(pos, currentObj, size, id, team));
}

void Unit::Draw() {
  this->MoveUnitPerFrame();
  if (this->data->IsSelected()) {
    DrawRectangleRec(this->data->GetRec(), GREEN);
  } else {
    DrawRectangleRec(this->data->GetRec(), this->data->color);
  }
};

ScreenElementType Unit::GetType() { return ScreenElementType::Unit; }

void Unit::MoveUnitPerFrame() {
  if (this->data->direction_vector_normalize.get() != nullptr) {
    // std::printf("CURRENT POS: X = %f, Y = %f \n", this->data->current_pos.x,
    //             this->data->current_pos.y);
    // std::printf("CURRENT OBJ: X = %f, Y = %f \n",
    // this->data->current_objective->x,
    //             this->data->current_objective->y);
    if ((this->data->current_pos.x == this->data->current_objective->x &&
         this->data->current_pos.y == this->data->current_objective->y) ||
        (this->data->current_objective->x < 0 &&
         this->data->current_objective->y < 0)) {
      return;
    }
    const float speed = 1.0;
    float moveX = this->data->direction_vector_normalize->x * speed;
    float moveY = this->data->direction_vector_normalize->y * speed;

    Vector2 toTarget = Vector2{
        .x = this->data->current_objective->x - this->data->current_pos.x,
        .y = this->data->current_objective->y - this->data->current_pos.y,
    };

    // std::cout << "MOVEMENT DES PIECES AVANT IF COLLIDE\n";
    // std::cout << this->data->current_pos.x << std::endl;
    // std::cout << this->data->current_pos.y << std::endl;
    if ((moveX * toTarget.x + moveY * toTarget.y) <= 0) {
      moveX = toTarget.x;
      moveY = toTarget.y;
    }

    this->data->current_pos.x += (float)moveX;
    this->data->current_pos.y += (float)moveY;

    this->data->rectangle.x = this->data->current_pos.x;
    this->data->rectangle.y = this->data->current_pos.y;

    // std::cout << "MOVEMENT DES PIECES\n";
    // std::cout << this->data->current_pos.x << std::endl;
    // std::cout << this->data->current_pos.y << std::endl;
    // std::printf("VALEUR DE CURRENT OBJ: x = %f, y = %f\n",
    //             this->data->GetCurrentObjective()->x,
    //             this->data->GetCurrentObjective()->y);
    // std::printf("VALEUR DE MOUVEMENT: x += %d, y += %d\n", moveX, moveY);
  };
}

// Building::Building(Vector2 pos, Vector2 currentObj, Vector2 size,
//                    std::string id, int team) {
//   this->data = std::make_shared<ElementData>(
//       ElementData(pos, currentObj, size, id, team));
// }

ElementData::ElementData(Rectangle rec) {
  this->rectangle = rec;
  this->is_selected = false;
};

std::shared_ptr<ElementData> Unit::GetElementData() { return this->data; };

ElementData::ElementData(Vector2 pos, Vector2 currentObj, Vector2 size,
                         std::string id, int team) {
  this->current_objective = std::make_shared<Vector2>(currentObj);
  this->current_pos = pos;
  this->rectangle = Rectangle{
      .x = pos.x,
      .y = pos.y,
      .width = size.x,
      .height = size.y,
  };
  this->id = id;
  this->is_selected = false;
  this->team = team;
  this->color = COLOR_TO_TEAM_ID.find(team)->second;
};

ElementData::~ElementData() {}

Vector2 ElementData::GetPos() { return this->current_pos; };

Rectangle ElementData::GetRec() { return this->rectangle; };

bool ElementData::IsSelected() { return this->is_selected; };

int ElementData::GetTeam() { return this->team; };

void ElementData::SetSelected(bool is_selected) {
  this->is_selected = is_selected;
};

void ElementData::SetPos(Vector2 vec2) {
  this->current_pos = vec2;
  // this->rectangle.x = vec2.x * SCALE_FACTOR;
  // this->rectangle.y = vec2.y * SCALE_FACTOR;
  this->rectangle.x = vec2.x;
  this->rectangle.y = vec2.y;
};

void ElementData::SetCurrentObjective(Vector2 vec2) {
  this->current_objective = std::make_shared<Vector2>(vec2);
};

void ElementData::SetCurrentObjective(std::shared_ptr<Vector2> vec2) {
  this->current_objective = vec2;
};

std::string ElementData::GetID() { return this->id; }

std::shared_ptr<Vector2> ElementData::GetCurrentObjective() {
  return this->current_objective;
}

void ElementData::SetDirectionVector() {
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
