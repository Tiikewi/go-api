package api

import (
	"encoding/json"
	"go-api/pkg/types"
	"net/http"
)

// @Summary Get ponged back.
// @Description getPing returns json with message key.
// @Tags Ping
// @Produce json
// @Router /ping [get]
// @Success 200 {object} types.PingResponse
// @Failure 404 {object} types.ErrorResponse
func getPing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

<<<<<<< Updated upstream
	var body types.PingResponse
	body.Message = "pong"
=======
	response := types.PingResponse{
		Message: db.GetVersion(),
	}
	if response.Message == "" {
		json.NewEncoder(w).Encode(
			types.ErrorResponse{
				StatusCode: 404,
				Message:    "Version not found",
			})
	}
>>>>>>> Stashed changes

	json.NewEncoder(w).Encode(response)
}
