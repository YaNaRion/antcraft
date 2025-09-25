#include "websocket.h"
#include "event_name.pb.h"
#include "raylib.h"
#include <iostream>

using websocketpp::lib::bind;
using websocketpp::lib::placeholders::_1;
using websocketpp::lib::placeholders::_2;

// SET URL FOR THE SERVER
const std::string url = "http://localhost:3000/ws";

void Gateway::OnMessage(websocketpp::connection_hdl hdl,
                        client::message_ptr msg) {
  try {
    if (msg->get_opcode() != websocketpp::frame::opcode::binary) {
      std::cerr << "Received non-binary frame; ignoring" << std::endl;
      return;
    }

    Event::Event event;
    if (!event.ParseFromString(msg->get_payload())) {
      std::cerr << "Failed to parse protobuf message!" << std::endl;
      return;
    }

    switch (event.Data_Event_case()) {
    case Event::Event::kMoveElement:
      if (event.has_move_element() && event.move_element().has_pos()) {
        const auto &units = event.move_element();
        std::cout << "NEW POS: X:" << units.pos().xpos()
                  << ", Y:" << units.pos().ypos() << std::endl;
        MoveUnitIN unitEvent = MoveUnitIN(
            Vector2{
                .x = (float)units.pos().xpos(),
                .y = (float)units.pos().ypos(),
            },
            Vector2{
                .x = (float)units.old_pos().xpos(),
                .y = (float)units.old_pos().ypos(),
            },
            units.player_id(), units.unit_id());
        std::shared_ptr<MoveUnitIN> shared_unit =
            std::make_shared<MoveUnitIN>(unitEvent);

        this->queue_in.push(shared_unit);
      }
      break;

    case Event::Event::kGameRespond:
      if (event.has_game_respond() && event.game_respond().has_map()) {
        const auto &game = event.game_respond();
        for (auto element : game.map().elements()) {

          Rectangle rec = Rectangle{
              .x = (float)element.element().pos().xpos(),
              .y = (float)element.element().pos().ypos(),
              .width = (float)element.element().size().xpos(),
              .height = (float)element.element().size().ypos(),
          };

          AddUnitIN unit = AddUnitIN(rec);
          std::shared_ptr<AddUnitIN> shared_unit =
              std::make_shared<AddUnitIN>(unit);
          this->queue_in.push(shared_unit);
        }
      }
      break;

    default:
      std::cerr << "Unknown event type" << std::endl;
    }
  } catch (websocketpp::exception const &e) {
    std::cerr << "WebSocket exception in OnMessage: " << e.what() << std::endl;
  }
}

Gateway::Gateway() {
  try {
    c.set_access_channels(websocketpp::log::alevel::all);
    c.clear_access_channels(websocketpp::log::alevel::frame_payload);
    c.init_asio();
    c.set_open_handler([this](websocketpp::connection_hdl hdl) {
      this->hdl = hdl;
      std::cout << "Connected to server!" << std::endl;
    });
    c.set_message_handler(bind(&Gateway::OnMessage, this, ::_1, ::_2));
    websocketpp::lib::error_code ec;
    client::connection_ptr con =
        c.get_connection("http://localhost:3000/ws", ec);
    std::cout << "APRES CONNECTION\n";
    con->replace_header("Origin", "http://localhost:3000");
    if (ec) {
      std::cout << "could not create connection because: " << ec.message()
                << std::endl;
    }
    c.connect(con);
    std::thread([this]() { c.run(); }).detach();
  } catch (websocketpp::exception const &e) {
    std::cout << e.what() << std::endl;
  }
};

Gateway::~Gateway() {};

void Gateway::PushEvent(std::shared_ptr<IEventOut> ev) {
  this->queue_out.push(ev);
};

size_t Gateway::GetQueueOutSize() { return this->queue_out.size(); }

void Gateway::PopAndSendEvent() {
  auto event = this->queue_out.front();
  Event::Event protoEvent = event->CreateProtoEvent();
  std::string eventString = protoEvent.SerializeAsString();
  std::cout << eventString << std::endl;
  this->SendEvent(eventString);
  this->queue_out.pop();
};

void Gateway::SendEvent(std::string eventString) {
  websocketpp::lib::error_code ec;
  this->c.send(this->hdl, eventString, websocketpp::frame::opcode::BINARY, ec);
  if (ec) {
    std::cout << "Send Error: " << ec.message() << std::endl;
  }
}

MoveUnitOut::MoveUnitOut(Vector2 old_p, Vector2 new_p, std::string player_ID,
                         std::string unit_ID) {
  old_pos = old_p;
  new_pos = new_p;
  playerID = player_ID;
  unitID = unit_ID;
};

Event::Event MoveUnitOut::CreateProtoEvent() {
  websocketpp::lib::error_code ec;
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

void MoveUnitIN::HandlerEvent(std::shared_ptr<GameScene> gameScene) {
  std::shared_ptr<IScreenElement> unit =
      gameScene->players.front()->GetUnit().front();
  unit->SetPos(this->new_pos);
  std::cout << "NOUVELLE POSITION RECU DU SERVER ET RECUP DANS CLEAN EVENT "
               "IN QUEUE \t X: "
            << unit->GetPos().x << "Y: " << unit->GetPos().y << std::endl;
  gameScene->AddElement(unit);
}

MoveUnitIN::MoveUnitIN(Vector2 old_p, Vector2 new_p, std::string player_ID,
                       std::string unit_ID) {
  old_pos = old_p;
  new_pos = new_p;
  playerID = player_ID;
  unitID = unit_ID;
};

AddUnitIN::AddUnitIN(Rectangle rec) { this->rec = rec; };

void AddUnitIN::HandlerEvent(std::shared_ptr<GameScene> gameScene) {
  Unit u = Unit(this->rec);
  std::shared_ptr<Unit> unit_share = std::make_shared<Unit>(u);
  gameScene->AddElement(unit_share);
}

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
