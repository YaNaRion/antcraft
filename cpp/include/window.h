#include "scene.h"

#include <memory>
#include <vector>

class Window {
public:
  Window(int width, int height, const char *title);
  void CloseWin();
  bool ShouldWindowClose();
  void Draw();

private:
  int width;
  int height;
  const char *title;
};
