#include "../include/raylib-cpp.hpp"
#include "window/window.h"

int main() {
	C_Window window = C_Window(800, 450);

	raylib::Color textColor(LIGHTGRAY);
	raylib::Window w(window.width_, window.height_, "Raylib C++ Starter Kit Example");

	SetTargetFPS(60);

	while (!w.ShouldClose()) {
		BeginDrawing();
		ClearBackground(RAYWHITE);
		textColor.DrawText("Congrats! You created your first window!", 190, 200, 20);
		EndDrawing();
	}

	return 0;
}
