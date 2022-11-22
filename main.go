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

	s := api.CreateNewServer()

	db := db.ConnectToDB()
	defer db.Close()

	fmt.Println("Server starting on port", port)
	http.ListenAndServe(port, s.Router)

}
