package app

type ErrorResponse struct {
	Message string `json:"error"`
	Code    int    `json:"code,omitempty"`
}
