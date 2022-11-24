package api

import (
	"encoding/json"
	"go-api/pkg/db"
	"go-api/pkg/types"
	"net/http"
)

// @Summary Get ponged back.
// @Description getPing returns json with message key.
// @Tags Ping
// @Produce json
// @Router /ping [get]
// @Success 200 {object} types.PingResponse
func (s *Server) getPing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	// Return status 200
	w.WriteHeader(http.StatusOK)

	var body types.PingResponse
	res := db.GetVersion(s.DB)
	body.Message = res

	json.NewEncoder(w).Encode(body)
}
