package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func PingHandler(r chi.Router) {
	r.Get("/", pingGET)
}

// Ping api Handler
func pingGET(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
