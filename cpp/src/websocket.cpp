#include "websocket.h"
#include "event_name.pb.h"
#include "raylib.h"
#include "unit.h"
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
    case Event::Event::kMoveUnit:
      if (event.has_move_unit() && event.move_unit().has_new_pos()) {
        const auto &units = event.move_unit();
        std::cout << "NEW POS: X:" << units.new_pos().xpos()
                  << ", Y:" << units.new_pos().ypos() << std::endl;
        MoveUnitPost unitEvent = MoveUnitPost(
            Vector2{
                .x = (float)units.new_pos().xpos(),
                .y = (float)units.new_pos().ypos(),
            },
            Vector2{
                .x = (float)units.old_pos().xpos(),
                .y = (float)units.old_pos().ypos(),
            },
            units.player_id(), units.unit_id());
        std::shared_ptr<MoveUnitPost> shared_unit =
            std::make_shared<MoveUnitPost>(unitEvent);

        this->queue_in.push(shared_unit);
      }
      break;

    case Event::Event::kPlayerData:
      // handle player data
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
    // Set logging to be pretty verbose (everything except message payloads)
    c.set_access_channels(websocketpp::log::alevel::all);
    c.clear_access_channels(websocketpp::log::alevel::frame_payload);

    c.init_asio();
    c.set_open_handler([this](websocketpp::connection_hdl hdl) {
      this->hdl = hdl;
      std::cout << "Connected to server!" << std::endl;
    });
    //
    // Register our message handler
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

    // Note that connect here only requests a connection. No network messages
    // are exchanged until the event loop starts running in the next line.
    c.connect(con);

    // Start the ASIO io_service run loop
    // this will cause a single connection to be made to the server. c.run()
    // will exit when this connection is closed.
    std::thread([this]() { c.run(); }).detach();

  } catch (websocketpp::exception const &e) {
    std::cout << e.what() << std::endl;
    std::cout << "DANS EXECEP INIT SOCKET\n";
  }
};

Gateway::~Gateway() {};

void Gateway::PushEvent(std::shared_ptr<IEvent> ev) {
  this->queue_out.push(ev);
};

size_t Gateway::GetQueueSize() { return this->queue_out.size(); }

void Gateway::PopEvent() {
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

MoveUnitPost::MoveUnitPost(Vector2 old_p, Vector2 new_p, std::string player_ID,
                           std::string unit_ID) {
  old_pos = old_p;
  new_pos = new_p;
  playerID = player_ID;
  unitID = unit_ID;
};

Event::Event MoveUnitPost::CreateProtoEvent() {
  websocketpp::lib::error_code ec;
  Event::Event event = Event::Event();

  Event::MoveUnit *unit = event.mutable_move_unit();
  unit->set_unit_id(this->unitID);
  unit->set_player_id(this->playerID);

  Event::Vector2 *old_pos = unit->mutable_old_pos();
  old_pos->set_xpos(this->old_pos.x);
  old_pos->set_ypos(this->old_pos.y);

  Event::Vector2 *new_pos = unit->mutable_new_pos();
  new_pos->set_xpos(this->new_pos.x);
  new_pos->set_ypos(this->new_pos.y);

  return event;
}

Vector2 MoveUnitPost::GetNewPos() { return this->new_pos; }
