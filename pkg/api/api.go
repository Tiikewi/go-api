package api

import (
	"github.com/go-chi/chi/v5"

	_ "go-api/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	Router *chi.Mux
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()

	s.MountHandlers()

	return s
}

func (s *Server) MountHandlers() {
	// Mount all Middleware here
	//s.Router.Use(middleware.Logger)

	// Swagger documentation
	s.Router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json")))

	// Mount all handlers here

	s.Router.Route("/ping", func(r chi.Router) {
		r.Get("/", getPing)
	})
}
