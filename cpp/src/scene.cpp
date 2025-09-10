#include "scene.h"
#include "raylib.h"
#include <iostream>

void MenuScene::Draw() {
  std::string menu_scene = "VOICI LA SCENE DU MENU";
  int fontSize = 40;
  DrawText(menu_scene.c_str(),
           (GetScreenWidth() / 2) - (menu_scene.size() * 10),
           GetScreenHeight() / 2 - (fontSize / 2), 40, RED);
}

MenuScene::MenuScene() { std::cout << "MENU SCENE\n"; };

MenuScene::~MenuScene() {};

SceneManager::SceneManager(std::vector<std::shared_ptr<IScene>> scenes) {
  this->scenes = scenes;
  // TODO Faire meilleur facon de faire le choix de scene
  if (this->scenes.size() > 1) {
    this->current_scene = scenes[1];
  } else if (this->scenes.size() > 0) {
    this->current_scene = scenes[0];
  }
}

SceneManager::~SceneManager() {}

void SceneManager::Draw() { this->current_scene->Draw(); }
