package api

import (
	"net/http"
)

func pingGET(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
