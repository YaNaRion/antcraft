#include "websocket.h"
#include "./event_name.pb.h"
#include <iostream>

// enum DataEventCase {
//   kPlayerData = 1,
//   kJoinRoomRequest = 2,
//   kJoinRoomResponse = 3,
//   kRoomStatusRequest = 4,
//   DATA_EVENT_NOT_SET = 0,
// };
using websocketpp::lib::bind;
using websocketpp::lib::placeholders::_1;
using websocketpp::lib::placeholders::_2;

// SET URL FOR THE SERVER
const std::string url = "http://localhost:3000/ws";
void WebsocketConnection::OnMessage(websocketpp::connection_hdl hdl,
                                    client::message_ptr msg) {
  if (this->hdl != nullptr) {
    this->hdl = &hdl;
  }

  Event::Event eventData;
  eventData.ParseFromString(msg->get_payload());

  // switch (eventData.Data_Event_case()) {
  // case Event::Event::kPlayerData:
  //   std::cout << "DANS CASE PLAYERDATA\n";
  //   Event::JoinRoomRequest player = eventData.join_room_request();
  //   std::cout << player.room() << "\n";
  //   break;
  // }
}

// void WebsocketConnection::OnMessage(websocketpp::connection_hdl hdl,
//                                     client::message_ptr msg) {
//
//   // std::cout << "on_message called with hdl: " << hdl.lock().get()
//   //           << " and message: " << msg->get_payload() << std::endl;
//   Event::Player player;
//   player.ParseFromString(msg->get_payload());
//   //
//   // player.ParseFromArray(msg->get_raw_payload(),
//   msg->get_payload().size()); std::cout << player.uniqueid() << "\n";
//
//   Event::Player playerResponde;
//   playerResponde.set_uniqueid("ECHO CLIENT FROM PROTO");
//
//   websocketpp::lib::error_code ec;
//   this->c.send(hdl, playerresponde.serializeasstring(), msg->get_opcode(),
//   ec);
//
//   if (ec) {
//     std::cout << "Echo failed because: " << ec.message() << std::endl;
//   }
// }

WebsocketConnection::WebsocketConnection() {
  try {
    // Set logging to be pretty verbose (everything except message payloads)
    c.set_access_channels(websocketpp::log::alevel::all);
    c.clear_access_channels(websocketpp::log::alevel::frame_payload);

    c.init_asio();

    // Register our message handler
    c.set_message_handler(
        bind(&WebsocketConnection::OnMessage, this, ::_1, ::_2));

  } catch (websocketpp::exception const &e) {
    std::cout << e.what() << std::endl;
  }
};

Gateway::~Gateway() {};

void Gateway::PushEvent(std::shared_ptr<IEvent> ev) { this->queue.push(ev); }

size_t Gateway::PopEvent() {
  if (this->queue.size() == 0) {
    return 0;
  }
  auto event = this->queue.front();
  event->PostEvent(this->ws.GetClient(), this->ws.GetHDL());
  this->queue.pop();
  return 1;
}

MoveUnit::MoveUnit(Vector2 old_p, Vector2 new_p, std::string player_ID,
                   std::string unit_ID) {
  old_pos = old_p;
  new_pos = new_p;
  playerID = player_ID;
  unitID = unit_ID;
};

void MoveUnit::PostEvent(client *c, websocketpp::connection_hdl *hdl) {
  websocketpp::lib::error_code ec;
  std::cout << "DANS POST EVENT MOVE UNIT\n";
  c->send(hdl, "payload", NULL, ec);
}
