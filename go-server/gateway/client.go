package gateway

import (
	"net"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

type ClientID uuid.UUID

type Client struct {
	Conn *websocket.Conn
	id   ClientID
	addr net.Addr
}

// Client that define the incomming connection, it is different for the playerID
func newClient(conn *websocket.Conn, id ClientID, addr net.Addr) *Client {
	return &Client{
		Conn: conn,
		id:   id,
		addr: addr,
	}
}

func (c *Client) GetID() ClientID {
	return c.id
}
