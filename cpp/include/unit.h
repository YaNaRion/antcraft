#include "raylib.h"

enum class ScreenElementType {
  Unit,
};

class IScreenElement {
public:
  virtual void Draw(Color color) = 0;
  virtual ~IScreenElement() {};
  virtual Vector2 GetPos() = 0;
  virtual ScreenElementType GetType() = 0;
};

class Unit : public IScreenElement {
public:
  Unit(Rectangle rec);
  ~Unit();
  void Draw(Color color) override;
  Vector2 GetPos() override;
  ScreenElementType GetType() override { return ScreenElementType::Unit; };

private:
  Rectangle rectangle;
  Vector2 *currentObjective;
};
