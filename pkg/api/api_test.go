package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-api/pkg/db"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"testing"
)

var s *Server

var DB *sql.DB

func setup() {
	port := ":" + os.Getenv("APP_PORT")
	// If no PORT variable given, default 8080.
	if port == ":" {
		port = ":8080"
	}

	DB = db.ConnectToDB()

	s = CreateNewServer(DB)

	fmt.Println("Server starting on port", port)
	http.ListenAndServe(port, s.Router)
}

func TestPing(t *testing.T) {
	os.Setenv("ENVIROMENT", "test")
	defer os.Unsetenv("ENVIROMENT")
	setup()
	defer DB.Close()

	//req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	//w := httptest.NewRecorder()

	const expectedMsg = "10.9.3-MariaDB-1:10.9.3+maria~ubu2204"

	l, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		http.HandleFunc("/ping", s.getPing)
		log.Fatal(http.Serve(l, nil))
	}()

	resp, err := http.Get("http://localhost:8082/ping")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error("Error when reading response body. Error:", err)
	}

	response := make(map[string]string)
	json.Unmarshal(body, &response)

	if response["message"] != expectedMsg {
		t.Error("Expected message: ", expectedMsg, ", Got", response["message"])
	}

}

//
