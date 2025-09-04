#include "window.h"
#include "raylib.h"

Window::Window(int width, int height, const char *title) {
  this->width = width;
  this->height = height;
  this->title = title;

  InitWindow(width, height, title);
}

void Window::CloseWin() { CloseWindow(); }
bool Window::ShouldWindowClose() { return WindowShouldClose(); }
