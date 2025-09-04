#include "raylib.h"
#include "window.h"
#include <memory>
#include <vector>

int main() {

  // Create MenuScene
  MenuScene menu = MenuScene();
  std::shared_ptr<MenuScene> menu_shared = std::make_shared<MenuScene>(menu);
  std::vector<std::shared_ptr<IScene>> vectorScene;

  vectorScene.push_back(menu_shared);

  Window window = Window(1920, 1080, "WINDOW FROM SCENE_MANAGER");
  SceneManager scene_manager = SceneManager(vectorScene);

  while (!window.ShouldWindowClose()) {
    BeginDrawing();
    ClearBackground(BLACK);

    scene_manager.Draw();

    EndDrawing();
  }

  window.CloseWin();

  return 0;
}
