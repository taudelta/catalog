package handler

import "github.com/gorilla/mux"

type Handler struct {
}

var handler *Handler

func init() {
	handler = &Handler{}
}

func RegisterEndpoints(router *mux.Router) {

	router.HandleFunc("/healthcheck", handler.Healtcheck).Methods("GET")

}
