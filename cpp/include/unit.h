#include "raylib.h"

class IUnit {
public:
  virtual void Draw(Color color) = 0;
  // virtual void ChangeObjective(Vector2 vec);
  virtual ~IUnit() {};
  // virtual void Move();
};

class Unit : public IUnit {
public:
  Unit(Vector2 pos);
  ~Unit();
  void Draw(Color color) override;
  // void ChangeObjective(Vector2 vec) override;
  // void Move() override;

private:
  Vector2 position;
  Vector2 *currentObjective;
};
