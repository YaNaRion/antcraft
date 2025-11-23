#include "event.h"
#include "event_name.pb.h"

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
                                 int team) {
  this->game_id = game_id;
  this->player_id = player_id;
  this->team = team;
}

void JoinGameEventIN::HandlerEvent(std::shared_ptr<GameScene> gameScene) {
  std::map<std::string, std::shared_ptr<Player>> players =
      gameScene->GetPlayers();
  Player player = Player(this->player_id, this->team);
  std::shared_ptr<Player> shared_player = std::make_shared<Player>(player);
  players.insert({player_id, shared_player});
  gameScene->SetCurrentPlayer(shared_player);
  return;
}

StartGameOut::StartGameOut(std::string gameID) { this->game_id = gameID; };

Event::Event StartGameOut::CreateProtoEvent() {
  Event::Event event = Event::Event();

  Event::StartGame *start_game = event.mutable_start_game();
  start_game->set_game_id(this->game_id);
  return event;
};

AttackUnit::AttackUnit(std::string gameID,
                       std::shared_ptr<IScreenElement> attaking_element,
                       std::shared_ptr<IScreenElement> attacked_element,
                       std::string playerID, int team) {
  this->gameID = gameID;
  this->attacked_element = attacked_element;
  this->attaking_element = attaking_element;
  this->playerID = playerID;
  this->team = team;
}

Event::Event AttackUnit::CreateProtoEvent() {
  Event::Event event = Event::Event();
  event.set_game_id(this->gameID);

  Event::Player *player = event.mutable_player_info();
  player->set_player_id(this->playerID);

  // A noter que cest pas la best facon de faire
  player->set_color(Event::ColorTeam(this->team));

  Event::AttackElement *attack_event = event.mutable_attack_element();

  attack_event->set_attacked_element_id(
      this->attacked_element->GetElementData()->GetID());

  attack_event->set_attacking_element_id(
      this->attaking_element->GetElementData()->GetID());

  return event;
};
