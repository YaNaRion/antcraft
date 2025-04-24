package main

import (
	"log"
	"main/controller"
	"main/gateway"
	"main/infra"
	"main/router"
	"main/service/config"
	"net/http"
	"time"
)

const (
	httpPort = ":3000"
)

// Write a fonction that print a chrisma tree

func Setup() *config.Server {
	// log.Println("Setup DB connection")
	var err error

	var db *infra.DB
	for db == nil {
		db, _ = infra.Setup()
		time.Sleep(1 * time.Second)
	}

	// Setup des routes de l'API
	log.Println("Setup Http controller")
	mux := http.NewServeMux()
	control := controller.SetUpController(mux, db)

	// Setup HTTP request
	log.Println("Setup Web router")
	router := router.Setup(mux)

	// Setup SocketIO connection
	log.Println("Setup SocketIO connection")
	socket, err := gateway.Setup(router.Mux, db)
	if err != nil {
		log.Println(err)
	}

	configServer := config.NewConf(db, router, socket, control)
	return config.NewServer(configServer)
}

func main() {
	// Setup DB connection
	server := Setup()
	log.Println("Listen on localhost:3000")
	err := http.ListenAndServe(httpPort, server.Conf.Router.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
