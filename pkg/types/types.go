package types

type PingResponse struct {
	Message string `json:"message"`
}

type RequestError struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}
