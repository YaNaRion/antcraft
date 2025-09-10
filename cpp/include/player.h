#include "raylib.h"
#include "unit.h"
#include <memory>
#include <vector>

class Player {
public:
  void Draw();
  Player(Color color, std::vector<std::shared_ptr<IUnit>> units);
  ~Player();
  Color GetTeamColor();

private:
  std::vector<std::shared_ptr<IUnit>> units;
  Color color;
};
