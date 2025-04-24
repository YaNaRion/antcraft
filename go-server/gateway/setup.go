package gateway

import (
	socketio "github.com/doquangtan/socket.io/v4"
	"log"
	"main/infra"
	"main/service/socket"
	"net/http"
)

const (
	robotRoom  = "robot room"
	webRoom    = "web room"
	scriptRoom = "script room"
)

type Gateway struct {
	io       *socketio.Io
	sockServ socket.SocketServiceInterface // Use the interface here
}

/*
* Fonction de set up de la route HTTP utiliser par SocketIO
 */
func Setup(mux *http.ServeMux, db *infra.DB) (*Gateway, error) {
	io := socketio.New()
	var socketService = socket.NewSocketService(db)

	socket, err := factoryGateway(io, socketService)
	if err != nil {
		return nil, errSocketCannotBeCreated
	}

	io.OnConnection(socket.handleConnection)

	mux.Handle("/socket.io/", io.HttpHandler())
	return socket, nil
}

func factoryGateway(
	io *socketio.Io,
	sockServ *socket.SocketService,
) (*Gateway, error) {
	if io == nil {
		return nil, errIOIsNil
	}
	if sockServ == nil {
		return nil, errServiceIsNil
	}
	socket := newGateway(io, sockServ)
	return &socket, nil
}

func newGateway(io *socketio.Io, sockServ *socket.SocketService) Gateway {
	return Gateway{
		io:       io,
		sockServ: sockServ,
	}
}

/*
* Handler de la première connection d'un client
 */
func (g *Gateway) handleConnection(socket *socketio.Socket) {
	log.Printf("New Socket Connection, user id: %s", socket.Id)

	// Event pour identifier les différents sockets
	socket.On(EventDisconnect, g.handleDisconnection)
}

// génère un nom de format "robot {i}".
// Retourne le premier nom disponible dans la map robots
// Détecte si le socket qui se déconnecte est un robot
func (g *Gateway) handleDisconnection(event *socketio.EventPayload) {}
