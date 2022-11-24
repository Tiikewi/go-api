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

	db := db.ConnectToDB()
	defer db.Close()

	s := api.CreateNewServer(db)

	fmt.Println("Server starting on port", port)
	http.ListenAndServe("0.0.0.0:8080", s.Router)

}
