#include "raylib.h"
#include <memory>
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
  virtual std::shared_ptr<Vector2> GetCurrentObjective() = 0;
  virtual void SetPos(Vector2 vec2) = 0;
  virtual bool IsSelected() = 0;
  virtual void SetSelected(bool is_selected) = 0;
  virtual std::string GetID() = 0;
  virtual Rectangle GetRec() = 0;
  virtual ScreenElementType GetType() = 0;
  virtual void SetDirectionVector() = 0;
  virtual void SetCurrentObjective(Vector2 vec2) = 0;
  virtual void SetCurrentObjective(std::shared_ptr<Vector2> vec2) = 0;
  virtual void MoveUnitPerFrame() = 0;
};

class Unit : public IScreenElement {
public:
  Unit(Rectangle rec);
  Unit(Vector2 pos, Vector2 currentObj, std::string id);
  ~Unit();
  void Draw(Color color) override;
  Vector2 GetPos() override;
  void SetPos(Vector2 vec2) override;
  bool IsSelected() override;
  void SetSelected(bool is_selected) override;
  std::string GetID() override;
  Rectangle GetRec() override;
  ScreenElementType GetType() override { return ScreenElementType::Unit; };
  void SetDirectionVector() override;
  void SetCurrentObjective(Vector2 vec2) override;
  void SetCurrentObjective(std::shared_ptr<Vector2> vec2) override;
  void MoveUnitPerFrame() override;
  std::shared_ptr<Vector2> GetCurrentObjective() override;

private:
  std::string id;
  Vector2 current_pos;
  Rectangle rectangle;
  std::shared_ptr<Vector2> current_objective;
  std::shared_ptr<Vector2> direction_vector_normalize;
  bool is_selected;
};
