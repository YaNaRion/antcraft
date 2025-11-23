#include <raylib.h>

int main(void) {
  InitWindow(800, 600, "Ant Animation Example");

  Texture2D ant = LoadTexture("assets/ants3.jpg");

  // Sprite sheet info
  const int frameWidth = 32;
  const int frameHeight = 32;
  const int frameCount = 4;
  int currentFrame = 0;

  float frameTime = 0.33f; // Seconds per frame
  float timer = 0.0f;

  Vector2 position = {400, 300};

  SetTargetFPS(60);

  while (!WindowShouldClose()) {
    // --- Update ---
    timer += GetFrameTime();
    if (timer >= frameTime) {
      timer = 0.0f;
      currentFrame++;
      if (currentFrame >= frameCount)
        currentFrame = 0;
    }

    // Source rectangle (selects one frame)
    Rectangle src = {float(frameWidth * currentFrame), 0, frameWidth,
                     frameHeight};

    // Destination rectangle (where and how big to draw)
    Rectangle dst = {position.x, position.y,
                     frameWidth * 2, // Draw scaled up (optional)
                     frameHeight * 2};

    // --- Draw ---
    BeginDrawing();
    ClearBackground(RAYWHITE);

    BeginBlendMode(BLEND_ALPHA);
    DrawTexturePro(ant, src, dst, (Vector2){frameWidth, frameHeight}, 0, WHITE);
    EndBlendMode();

    EndDrawing();
  }

  UnloadTexture(ant);
  CloseWindow();

  return 0;
}
