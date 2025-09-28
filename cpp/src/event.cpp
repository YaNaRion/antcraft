#include "event.h"

MoveUnitOut::MoveUnitOut(Vector2 old_p, Vector2 new_p, std::string player_ID,
                         std::string unit_ID) {
  old_pos = old_p;
  new_pos = new_p;
  playerID = player_ID;
  unitID = unit_ID;
};

Event::Event MoveUnitOut::CreateProtoEvent() {
  Event::Event event = Event::Event();

  Event::MoveElement *unit = event.mutable_move_element();
  unit->set_unit_id(this->unitID);
  unit->set_player_id(this->playerID);

  Event::Vector2 *old_pos = unit->mutable_old_pos();
  old_pos->set_xpos(this->old_pos.x);
  old_pos->set_ypos(this->old_pos.y);

  Event::Vector2 *new_pos = unit->mutable_pos();
  new_pos->set_xpos(this->new_pos.x);
  new_pos->set_ypos(this->new_pos.y);

  return event;
}

Vector2 MoveUnitOut::GetNewPos() { return this->new_pos; }

GameResquest::GameResquest(std::string playerID, std::string gameID) {
  this->gameID = gameID;
  this->playerID = playerID;
}

Event::Event GameResquest::CreateProtoEvent() {
  Event::Event eventWrapper;
  Event::ConnectToGameRequest *event = eventWrapper.mutable_game_request();
  event->set_game_id(this->gameID);
  event->set_player_id(this->playerID);

  return eventWrapper;
}
UpdateMapStateEvent::UpdateMapStateEvent(
    std::vector<std::shared_ptr<IScreenElement>> elements,
    Event::GameState state) {
  this->elements = elements;
  this->state = state;
}

void UpdateMapStateEvent::HandlerEvent(std::shared_ptr<GameScene> game_scene) {
  std::cout << "DANS HANDLER UPDATE EVENT" << std::endl;
  game_scene->Update(this->elements, this->state);
  std::cout << "DANS HANDLER UPDATE EVENT" << std::endl;
}
