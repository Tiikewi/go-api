package main

import (
	"fmt"
	"go-api/handlers"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	port := ":8080"
	s := CreateNewServer()
	s.MountHandlers()
	fmt.Println("Server runnin on port", port)
	http.ListenAndServe(port, s.Router)
}

func (s *Server) MountHandlers() {
	// Mount all Middleware here
	s.Router.Use(middleware.Logger)

	// Mount all handlers here
	s.Router.Route("/ping", handlers.PingHandler)
}
