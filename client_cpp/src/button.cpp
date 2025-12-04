#include "button.h"
#include "raylib.h"

const std::string Play_Button_String = "./assets/play_button_3.png";

Button::Button(std::string message, Vector2 position) {
  this->texture = LoadTexture(Play_Button_String.c_str());
  this->position = position;

  this->position.x -= (float)this->texture.width / 2;
  this->position.y -= (float)this->texture.height / 2;

  this->message = message;
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

  int font_size = 36;
  int text_width = MeasureText(this->message.c_str(), font_size);

  Vector2 TextPosition = Vector2{
      .x = rec.x + rec.width / 2 - text_width / 2,
      .y = rec.y + rec.height / 2 - font_size / 2,
  };

  Color color = WHITE;
  if (CheckCollisionPointRec(GetMousePosition(), rec))
    color = PURPLE;

  DrawTexture(this->texture, this->position.x, this->position.y, color);
  DrawText(this->message.c_str(), TextPosition.x, TextPosition.y, font_size,
           color);
}
