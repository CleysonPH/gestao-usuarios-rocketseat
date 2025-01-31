package api

type response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}
