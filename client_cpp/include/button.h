#include "raylib.h"
#include <string>

class Button {
public:
  Button();
  Button(std::string message, Vector2 position);
  ~Button();
  Vector2 position;
  Texture2D texture;
  std::string message;

  Vector2 GetSize();
  Vector2 GetPos();
  Rectangle GetRect();
  void Draw();
};
