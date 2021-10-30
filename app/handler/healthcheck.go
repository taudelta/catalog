package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "ok")
}
