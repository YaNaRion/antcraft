#include "scene.h"
#include "raylib.h"
#include <iostream>

void MenuScene::Draw() {
  DrawText("VOICI LA SCENE DU MENU", 800, 400, 40, RED);
}

MenuScene::MenuScene() { std::cout << "MENU SCENE\n"; };

MenuScene::~MenuScene() {};

SceneManager::SceneManager(std::vector<std::shared_ptr<IScene>> scenes) {
  this->scenes = scenes;
  // TODO Faire meilleur facon de faire le choix de scene
  if (this->scenes.size() > 0) {
    this->current_scene = scenes[0];
  }
}

SceneManager::~SceneManager() {}

void SceneManager::Draw() { this->current_scene->Draw(); }
