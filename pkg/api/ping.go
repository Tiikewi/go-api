package api

import (
	"encoding/json"
	"go-api/pkg/db"
	"go-api/pkg/types"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func handlePing(r chi.Router) {
	r.Get("/", getPing)
}

// @Summary Get ponged back.
// @Description getPing returns json with message key.
// @Tags Ping
// @Produce json
// @Router /ping [get]
// @Success 200 {object} types.PingResponse
func getPing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	// Return status 200
	w.WriteHeader(http.StatusOK)

	var body types.PingResponse
	body.Message = db.GetVersion()

	json.NewEncoder(w).Encode(body)
}
