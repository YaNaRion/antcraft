#include "raylib.h"
#include "unit.h"
#include <memory>
#include <vector>

class Player {
public:
  void Draw();
  Player(Color color, std::vector<std::shared_ptr<IScreenElement>> units,
         std::string player_id);
  ~Player();
  Color GetTeamColor();
  void AddElement(std::shared_ptr<IScreenElement> elem);

  std::vector<std::shared_ptr<IScreenElement>> GetUnits();

private:
  std::string player_id;
  std::vector<std::shared_ptr<IScreenElement>> units;
  Color color;
};
