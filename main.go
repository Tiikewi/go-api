package main

import (
	"fmt"
	"go-api/pkg/api"
	"go-api/pkg/db"
	"net/http"
	"os"

	_ "go-api/docs"
)

// @title Golang example API
// @version 0.1
// @description This is a simple boilerplate for Go API.

// @contact.name Kimi Porthan

// @host localhost:8080

func main() {
	port := ":" + os.Getenv("APP_PORT")
	// If no PORT variable given, default 8080.
	if port == ":" {
		port = ":8080"
	}

	db.ConnectToDB()
	defer db.DB.Close()

	s := api.CreateNewServer()

	fmt.Println("Server running on port", port)
	http.ListenAndServe(port, s.Router)
}
