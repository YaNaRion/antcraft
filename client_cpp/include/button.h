#include "raylib.h"

class Button {
public:
  Button();
  Button(Vector2 position);
  ~Button();
  Vector2 position;
  Texture2D texture;

  Vector2 GetSize();
  Vector2 GetPos();
  Rectangle GetRect();
  void Draw();
};
