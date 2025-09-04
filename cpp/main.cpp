#include "raylib.h"
#include "scene.h"
#include <memory.h>
#include <vector>

int main() {
  InitWindow(800, 450, "raylib [core] example - basic window");
  MenuScene menu = MenuScene();

  std::vector<std::shared_ptr<MenuScene>> vectorScene;
  vectorScene.push_back(menu);

  while (!WindowShouldClose()) {
    BeginDrawing();
    ClearBackground(RAYWHITE);
    DrawText("Congrats! You created your first window!", 190, 200, 20,
             LIGHTGRAY);
    EndDrawing();
  }

  CloseWindow();
  return 0;
}
