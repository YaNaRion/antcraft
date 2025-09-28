#include "raylib.h"
#include <string>

const int SCALE_FACTOR = 10;

enum class ScreenElementType {
  Unit,
};

class IScreenElement {
public:
  virtual ~IScreenElement() {};
  virtual void Draw(Color color) = 0;
  virtual Vector2 GetPos() = 0;
  virtual void SetPos(Vector2 vec2) = 0;
  virtual bool IsSelected() = 0;
  virtual void SetSelected(bool is_selected) = 0;
  virtual std::string GetID() = 0;
  virtual Rectangle GetRec() = 0;
  virtual ScreenElementType GetType() = 0;
};

class Unit : public IScreenElement {
public:
  Unit(Rectangle rec);
  Unit(Vector2 pos, std::string id);
  ~Unit();
  void Draw(Color color) override;
  Vector2 GetPos() override;
  void SetPos(Vector2 vec2) override;
  bool IsSelected() override;
  void SetSelected(bool is_selected) override;
  std::string GetID() override;
  Rectangle GetRec() override;
  ScreenElementType GetType() override { return ScreenElementType::Unit; };

private:
  std::string id;
  Vector2 current_pos;
  Rectangle rectangle;
  Vector2 *current_objective;
  bool is_selected;
};
