package main

import (
	"log"
	"main/controller"
	"main/gateway"
	"main/infra"
	"main/router"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

const (
	httpPort = ":3000"
)

type Config struct {
	Router     *router.Router
	DB         *infra.DB
	Controller *controller.Controller
	Websock    *gateway.WebsocketManager
}

func NewConf(
	db *infra.DB,
	router *router.Router,
	controller *controller.Controller,
	websocket *gateway.WebsocketManager,
) *Config {
	return &Config{DB: db, Router: router, Controller: controller, Websock: websocket}
}

type Server struct {
	Conf *Config
}

func NewServer(config *Config) *Server {
	return &Server{
		Conf: config,
	}
}

// Write a fonction that print a chrisma tree

func Setup(logger *log.Logger) *Server {
	// log.Println("Setup DB connection")
	var db *infra.DB

	// Setup des routes de l'API
	logger.Println("Setup Http controller")
	mux := http.NewServeMux()
	control := controller.SetUpController(mux, db, logger)

	// Setup HTTP request
	logger.Println("Setup Web router")
	router := router.Setup(mux, logger)

	// Setup SocketIO connection
	logger.Println("Setup SocketIO connection")
	websocket := gateway.NewWebsocketManager(logger)

	configServer := NewConf(db, router, control, websocket)
	return NewServer(configServer)
}

func main() {
	logger := log.New(os.Stdout, "LOG: ", log.LstdFlags)
	server := Setup(logger)
	log.Println("SERVER STARTED")
	server.Conf.Router.Mux.Handle("/ws", websocket.Handler(server.Conf.Websock.HandleWS))
	err := http.ListenAndServe(httpPort, server.Conf.Router.Mux)
	if err != nil {
		log.Println(err)
	}
}
