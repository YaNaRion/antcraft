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
  old_pos->set_x(this->old_pos.x);
  old_pos->set_y(this->old_pos.y);

  Event::Vector2 *new_pos = unit->mutable_new_pos();
  new_pos->set_x(this->new_pos.x);
  new_pos->set_y(this->new_pos.y);

  std::string gameID = "Game1";
  event.set_game_id(gameID);

  Event::Player *player_info = event.mutable_player_info();
  player_info->set_player_id(this->playerID);

  return event;
}

Vector2 MoveUnitOut::GetNewPos() { return this->new_pos; }

GameResquest::GameResquest(std::string playerID, std::string gameID) {
  this->gameID = gameID;
  this->playerID = playerID;
}

Event::Event GameResquest::CreateProtoEvent() {
  Event::Event eventWrapper;
  Event::JoinGame *event = eventWrapper.mutable_join_game();
  eventWrapper.set_game_id(this->gameID);
  Event::Player *player_info = eventWrapper.mutable_player_info();
  player_info->set_player_id(this->playerID);
  player_info->set_color(Event::RED_PROTO);
  return eventWrapper;
}

UpdateMapStateEventIN::UpdateMapStateEventIN(
    std::vector<std::shared_ptr<IScreenElement>> elements,
    Event::GameState state) {
  this->elements = elements;
  this->state = state;
}

void UpdateMapStateEventIN::HandlerEvent(
    std::shared_ptr<GameScene> game_scene) {
  game_scene->Update(this->elements, this->state);
}

JoinGameEventIN::JoinGameEventIN(std::string game_id, std::string player_id,
                                 Color color) {
  this->game_id = game_id;
  this->player_id = player_id;
  this->color = color;
}

void JoinGameEventIN::HandlerEvent(std::shared_ptr<GameScene> gameScene) {
  std::map<std::string, std::shared_ptr<Player>> players =
      gameScene->GetPlayers();
  Player player = Player(this->player_id, this->color);
  std::shared_ptr<Player> shared_player = std::make_shared<Player>(player);
  players.insert({player_id, shared_player});
  gameScene->SetCurrentPlayer(shared_player);
}
