#include "raylib.h"
#include <map>
#include <memory>
#include <string>

const std::map<int, Color> COLOR_TO_TEAM_ID = {{1, RED}, {2, BLUE}};
const int SCALE_FACTOR = 10;

enum class ScreenElementType {
  Unit,
  Base,
};

class ElementData {
public:
  ElementData(Rectangle rec);
  ElementData(Vector2 pos, Vector2 size, std::string id, int team);
  ~ElementData();

  Vector2 GetPos();
  void SetPos(Vector2 vec2);
  bool IsSelected();
  void SetSelected(bool is_selected);
  std::string GetID();
  Rectangle GetRec();
  void SetPosition();
  void SetDirectionVector();
  void SetCurrentObjective(Vector2 vec2);
  void SetCurrentObjective(std::shared_ptr<Vector2> vec2);
  void MoveUnitPerFrame();
  int GetTeam();
  std::shared_ptr<Vector2> GetCurrentObjective();

  friend class Unit;
  friend class Building;

private:
  std::string id;
  Vector2 current_pos;
  Rectangle rectangle;
  std::shared_ptr<Vector2> current_objective;
  std::shared_ptr<Vector2> direction_vector_normalize;
  bool is_selected;
  int team;
  Color color;
};

class IScreenElement {
public:
  virtual ~IScreenElement() {};
  virtual void Draw() = 0;
  virtual ScreenElementType GetType() = 0;
  virtual std::shared_ptr<ElementData> GetElementData() = 0;
  // virtual Vector2 GetPos() = 0;
  // virtual std::shared_ptr<Vector2> GetCurrentObjective() = 0;
  // virtual void SetPos(Vector2 vec2) = 0;
  // virtual bool IsSelected() = 0;
  // virtual void SetSelected(bool is_selected) = 0;
  // virtual std::string GetID() = 0;
  // virtual Rectangle GetRec() = 0;
  // virtual void SetDirectionVector() = 0;
  // virtual void SetCurrentObjective(Vector2 vec2) = 0;
  // virtual void SetCurrentObjective(std::shared_ptr<Vector2> vec2) = 0;
  // virtual void MoveUnitPerFrame() = 0;
  // virtual int GetTeam() = 0;
};

class Unit : public IScreenElement {
public:
  Unit(Vector2 pos, Vector2 size, std::string id, int team);
  ~Unit() {};
  void Draw() override;
  std::shared_ptr<ElementData> GetElementData() override;
  ScreenElementType GetType() override;
  void MoveUnitPerFrame();

private:
  std::shared_ptr<ElementData> data;
};

class Building : public IScreenElement {
public:
  Building(Vector2 pos, Vector2 size, std::string id, int team);
  ~Building() {};
  void Draw() override;
  ScreenElementType GetType() override;
  std::shared_ptr<ElementData> GetElementData() override;

private:
  std::shared_ptr<ElementData> data;
};
