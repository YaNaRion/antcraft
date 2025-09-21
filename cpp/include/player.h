#include "raylib.h"
#include "unit.h"
#include <memory>
#include <vector>

class Player {
public:
  void Draw();
  Player(Color color, std::vector<std::shared_ptr<IScreenElement>> units);
  ~Player();
  Color GetTeamColor();
  std::vector<std::shared_ptr<IScreenElement>> GetUnit();

private:
  std::vector<std::shared_ptr<IScreenElement>> units;
  Color color;
};
