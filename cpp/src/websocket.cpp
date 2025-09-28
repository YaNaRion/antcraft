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
  case Event::Event::kSyncGameState:
    if (event.has_sync_game_state()) {
      const auto &game_sync = event.sync_game_state();
      std::vector<std::shared_ptr<IScreenElement>> vector;
      std::cout << "DANS SYNC\n";
      game_sync.elements();

      for (const Event::Element &elem : game_sync.elements()) {
        std::cout << "DANS SYNC FOR\n";
        const auto element = elem.element();
        std::cout << element.unit_id() << std::endl;
        std::cout << "DANS SYNC FOR\n";
        Vector2 pos = Vector2{
            .x = (float)element.pos().xpos(),
            .y = (float)element.pos().ypos(),
        };

        Unit unit = Unit(pos, elem.element().unit_id());
        vector.push_back(std::make_shared<Unit>(unit));
      }
      std::cout << "DANS SYNC\n";
      UpdateMapStateEvent update_map_event =
          UpdateMapStateEvent(vector, game_sync.game_state());
      std::shared_ptr<UpdateMapStateEvent> share_update =
          std::make_shared<UpdateMapStateEvent>(update_map_event);
      this->queue_in.push(share_update);
    }
    break;
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
