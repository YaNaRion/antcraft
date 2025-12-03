#include "button.h"
#include "raylib.h"
#include <string>

const std::string Play_Button_String = "./assets/Playbutton2.png";

Button::Button(Vector2 position) {
  this->texture = LoadTexture(Play_Button_String.c_str());
  this->position = position;

  this->position.x -= (float)this->texture.width / 2;
  this->position.y -= (float)this->texture.height / 2;
}

Button::~Button() {}

Vector2 Button::GetPos() { return this->position; }

Rectangle Button::GetRect() {
  Vector2 size = this->GetSize();
  return Rectangle{
      .x = this->position.x,
      .y = this->position.y,
      .width = size.x,
      .height = size.y,
  };
}

Vector2 Button::GetSize() {
  return Vector2{
      .x = (float)this->texture.width,
      .y = (float)this->texture.height,
  };
}

void Button::Draw() {
  Rectangle rec = {
      .x = this->position.x,
      .y = this->position.y,
      .width = this->GetSize().x,
      .height = this->GetSize().y,
  };

  if (CheckCollisionPointRec(GetMousePosition(), rec)) {
    DrawTexture(this->texture, this->position.x, this->position.y, PURPLE);
  } else {
    DrawTexture(this->texture, this->position.x, this->position.y, WHITE);
  }
}
