package controller

import (
	"log"
	"main/infra"
	"net/http"
)

const ContentTypeJSON = "application/json"

type Controller struct {
	db  *infra.DB
	mux *http.ServeMux
	log *log.Logger
}

func newController(db *infra.DB, mux *http.ServeMux, log *log.Logger) *Controller {
	return &Controller{
		db:  db,
		mux: mux,
		log: log,
	}
}

func SetUpController(mux *http.ServeMux, db *infra.DB, log *log.Logger) *Controller {
	controller := newController(db, mux, log)
	mux.HandleFunc("GET /missions/", controller.getRandom)
	return controller
}

func writeResponse(w http.ResponseWriter, data []byte) {
	w.Header().Set("Contend-Type", ContentTypeJSON)
	_, err := w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
