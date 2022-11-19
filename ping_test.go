package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// executeRequest, creates a new ResponseRecorder
// then executes the request by calling ServeHTTP in the router
// after which the handler writes the response to the response recorder
// which we can then inspect.

func TestPing(t *testing.T) {
	// Create a New Server Struct
	s := CreateNewServer()
	// Mount Handlers
	s.MountHandlers()

	// Create a New Request
	req, _ := http.NewRequest("GET", "/ping", nil)

	// Execute Request
	response := executeRequest(req, s)

	// Check the response code
	checkResponseCode(t, http.StatusOK, response.Code)

	// We can use testify/require to assert values, as it is more convenient
	require.Equal(t, "pong", response.Body.String())
}
