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
  Event::Event eventData;
  eventData.ParseFromString(msg->get_payload());

  switch (eventData.Data_Event_case()) {
  case Event::Event::kPlayerData:
    std::cout << "DANS CASE PLAYERDATA\n";
    Event::JoinRoomRequest player = eventData.join_room_request();
    std::cout << player.room() << "\n";
    break;
  }
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
//   this->c.send(hdl, playerResponde.SerializeAsString(), msg->get_opcode(),
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

    websocketpp::lib::error_code ec;
    client::connection_ptr con = c.get_connection(url, ec);

    con->replace_header("Origin", "http://localhost:3000");

    if (ec) {
      std::cout << "could not create connection because: " << ec.message()
                << std::endl;
    }

    c.connect(con);
    c.run();
  } catch (websocketpp::exception const &e) {
    std::cout << e.what() << std::endl;
  }
};
