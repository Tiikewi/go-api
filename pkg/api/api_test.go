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
	"net/http/httptest"
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
	l, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		http.ListenAndServe(":8082", s.Router)
		log.Fatal(http.Serve(l, nil))
	}()
}

func TestPing(t *testing.T) {
	os.Setenv("ENVIROMENT", "test")
	defer os.Unsetenv("ENVIROMENT")
	setup()
	defer DB.Close()

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8082/ping", nil)
	w := httptest.NewRecorder()

	const expectedMsg = "10.9.3-MariaDB-1:10.9.3+maria~ubu2204"

	s.getPing(w, req)
	res := w.Result()
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
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
