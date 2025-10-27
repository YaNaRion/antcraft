package game

import "golang.org/x/net/websocket"

type PlayerConn struct {
	ws *websocket.Conn
}

func NewPlayerConn(ws *websocket.Conn) *PlayerConn {
	return &PlayerConn{ws: ws}
}

type PlayerID string

type Player struct {
	playerID PlayerID
	elements []IElement
	building []Building
	conn     *PlayerConn
}

func NewPlayer(conn *PlayerConn, playerID PlayerID) *Player {
	return &Player{
		elements: make([]IElement, 0),
		conn:     conn,
		playerID: playerID,
	}
}

func (p *Player) AddElement(el IElement) {
	p.elements = append(p.elements, el)
}

func (p *Player) AddBuilding(el Building) {
	p.building = append(p.building, el)
}

func (p *Player) GetPlayerIDString() string {
	return string(p.playerID)
}
