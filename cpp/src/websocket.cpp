#include "websocket.h"
#include "event_name.pb.h"
#include "raylib.h"
#include <boost/asio/serial_port_base.hpp>
#include <iostream>
#include <memory>

using websocketpp::lib::bind;
using websocketpp::lib::placeholders::_1;
using websocketpp::lib::placeholders::_2;

const std::string url = "http://localhost:3000/ws";

void Gateway::OnMessage(websocketpp::connection_hdl hdl,
                        client::message_ptr msg) {
  if (msg->get_opcode() != websocketpp::frame::opcode::binary) {
    std::cerr << "Received non-binary frame; ignoring" << std::endl;
    return;
  }

  Event::Event event;
  if (!event.ParseFromString(msg->get_payload())) {
    std::cerr << "Failed to parse protobuf message!" << std::endl;
    return;
  }

  switch (event.data_event_case()) {
  case Event::Event::kSyncGameState:
    if (event.has_sync_game_state()) {
      const auto &game_sync = event.sync_game_state();
      this->SyncGameEventHandler(game_sync);
    }
    break;
  case Event::Event::kJoinGame:
    const auto &game_sync = event.join_game();
    JoinGameEventIN join_game_event =
        JoinGameEventIN(event.game_id(), event.player_info().player_id(),
                        event.player_info().color());

    this->queue_in.push(std::make_shared<JoinGameEventIN>(join_game_event));
    break;
  }
}

Gateway::Gateway() {};

Gateway::~Gateway() {};

bool Gateway::GetIsConnected() { return this->IsConnected; }

void Gateway::Connect() {
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
    con->replace_header("Origin", "http://localhost:3000");
    if (ec) {
      std::cout << "could not create connection because: " << ec.message()
                << std::endl;
    }
    c.connect(con);
    std::thread([this]() { c.run(); }).detach();
    this->IsConnected = true;
  } catch (websocketpp::exception const &e) {
    std::cout << e.what() << std::endl;
  }
}

void Gateway::PushEvent(std::shared_ptr<IEventOut> ev) {
  this->queue_out.push(ev);
};

size_t Gateway::GetQueueOutSize() { return this->queue_out.size(); }

void Gateway::PopAndSendEvent() {
  // std::cout << "DANS POP AND SERVER" << std::endl;
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

void Gateway::SyncGameEventHandler(const Event::SyncGameState &game_state) {
  std::vector<std::shared_ptr<IScreenElement>> vector;

  for (const Event::Element &element : game_state.elements()) {
    float x_pos = (float)element.pos().x();
    float y_pos = (float)element.pos().y();
    Vector2 pos = Vector2{
        .x = x_pos,
        .y = y_pos,
    };

    float x_size = (float)element.size().x();
    float y_size = (float)element.size().y();
    Vector2 size = Vector2{
        .x = x_size,
        .y = y_size,
    };

    Vector2 currentObjectif = Vector2{
        .x = (float)element.currentobjective().x(),
        .y = (float)element.currentobjective().y(),
    };

    if (element.element_type() == Event::ElementType::WORKER) {
      Unit unit =
          Unit(pos, currentObjectif, size, element.unit_id(), element.team());
      std::shared_ptr<Unit> unit_shared = std::make_shared<Unit>(unit);
      vector.push_back(unit_shared);
    }

    if (element.element_type() == Event::ElementType::BASE) {
      Building building = Building(pos, currentObjectif, size,
                                   element.unit_id(), element.team());
      std::shared_ptr<Building> building_shared =
          std::make_shared<Building>(building);
      vector.push_back(building_shared);
    }
  }

  UpdateMapStateEventIN update_map_event =
      UpdateMapStateEventIN(vector, game_state.game_state());

  std::shared_ptr<UpdateMapStateEventIN> share_update =
      std::make_shared<UpdateMapStateEventIN>(update_map_event);
  this->queue_in.push(share_update);
};
