#include "scene.h"
#include "raylib.h"
#include <iostream>
#include <string>

const std::string Play_Button_String = "./assets/Playbutton2.png";

void MenuScene::Draw() {
  DrawTexture(this->PlayButtonTex, (GetScreenWidth() / 2) - 256 / 2,
              GetScreenHeight() / 2 - 62 / 2, WHITE);
}

MenuScene::MenuScene() {
  this->PlayButtonTex = LoadTexture(Play_Button_String.c_str());
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
