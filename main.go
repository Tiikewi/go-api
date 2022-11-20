package main

import (
	"fmt"
	"go-api/pkg/api"
	"net/http"
	"os"
)

func main() {
	port := ":" + os.Getenv("APP_PORT")

	s := api.CreateNewServer()
	s.MountHandlers()

	fmt.Println("Server starting on port", port)
	http.ListenAndServe(port, s.Router)
}
