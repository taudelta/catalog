package handler

import "github.com/gorilla/mux"

type Handler struct {
}

var handler *Handler

func init() {
	handler = &Handler{}
}

func RegisterEndpoints(router *mux.Router) {

	router.HandleFunc("/healthcheck", handler.Healthcheck).Methods("GET")

	router.HandleFunc("/schema", handler.CreateSchema).Methods("POST")
	router.HandleFunc("/table", handler.CreateTable).Methods("POST")

}
