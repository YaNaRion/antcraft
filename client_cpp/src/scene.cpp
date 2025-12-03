#include "scene.h"
#include "raylib.h"
#include <iostream>
#include <string>

const std::string backgound_texture_path = "./assets/menu_background.png";

void MenuScene::Draw() {

  DrawTexture(this->backgound_texture, 0, 0, WHITE);
  DrawText("DEBUG: MENU SCENE", 10, 10, 36, BLACK);
  this->play_button.Draw();
}

void MenuScene::SetBackgoundTexture() {
  this->backgound_texture = LoadTexture(backgound_texture_path.c_str());
};

MenuScene::~MenuScene() {};

SceneManager::SceneManager(std::shared_ptr<GameScene> game_scene,
                           std::shared_ptr<MenuScene> menu_scene) {
  this->game_scene = game_scene;
  this->menu_scene = menu_scene;
  this->current_scene = menu_scene;
}

SceneManager::~SceneManager() {}

void SceneManager::Draw() { this->current_scene->Draw(); }
