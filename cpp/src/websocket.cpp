#include "websocket.h"
#include <iostream>

using websocketpp::lib::bind;
using websocketpp::lib::placeholders::_1;
using websocketpp::lib::placeholders::_2;

void WebsocketConnection::OnMessage(websocketpp::connection_hdl hdl,
                                    client::message_ptr msg) {

  std::cout << "on_message called with hdl: " << hdl.lock().get()
            << " and message: " << msg->get_payload() << std::endl;

  websocketpp::lib::error_code ec;
  this->c.send(hdl, "ECHO CLIENT", msg->get_opcode(), ec);

  if (ec) {
    std::cout << "Echo failed because: " << ec.message() << std::endl;
  }
}

WebsocketConnection::WebsocketConnection() {
  try {
    // Set logging to be pretty verbose (everything except message payloads)

    c.set_access_channels(websocketpp::log::alevel::all);
    c.clear_access_channels(websocketpp::log::alevel::frame_payload);

    // Initialize ASIO
    c.init_asio();

    // Register our message handler
    c.set_message_handler(
        bind(&WebsocketConnection::OnMessage, this, ::_1, ::_2));

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
    c.run();
  } catch (websocketpp::exception const &e) {
    std::cout << e.what() << std::endl;
  }
};
