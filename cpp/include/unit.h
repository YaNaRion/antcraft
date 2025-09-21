#include "raylib.h"

enum class ScreenElementType {
  Unit,
};

class IScreenElement {
public:
  virtual void Draw(Color color) = 0;
  virtual ~IScreenElement() {};
  virtual Vector2 GetPos() = 0;

  virtual void SetPos(Vector2 vec2) = 0;
  virtual Rectangle GetRec() = 0;
  virtual bool IsSelected() = 0;
  virtual void SetSelected(bool is_selected) = 0;

  virtual ScreenElementType GetType() = 0;
};

class Unit : public IScreenElement {
public:
  Unit(Rectangle rec);
  ~Unit();
  void Draw(Color color) override;

  Vector2 GetPos() override;
  void SetPos(Vector2 vec2) override;

  bool IsSelected() override;
  void SetSelected(bool is_selected) override;
  Rectangle GetRec() override;
  ScreenElementType GetType() override { return ScreenElementType::Unit; };

private:
  Rectangle rectangle;
  Vector2 *current_objective;
  bool is_selected;
};
