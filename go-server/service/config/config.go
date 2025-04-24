package config

import (
	"main/controller"
	socket "main/gateway"
	"main/infra"
	"main/router"
)

type Config struct {
	Router     *router.Router
	DB         *infra.DB
	Socket     *socket.Gateway
	Controller *controller.Controller
}

func NewConf(
	db *infra.DB,
	router *router.Router,
	socket *socket.Gateway,
	controller *controller.Controller,
) *Config {
	return &Config{DB: db, Router: router, Socket: socket}
}
