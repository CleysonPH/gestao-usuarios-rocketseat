package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func sendJSON(w http.ResponseWriter, r response, status int) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(r)
	if err != nil {
		slog.Error("failed to marshal json data", "error", err)
		sendJSON(w, response{Error: "something went wrong"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write response to client", "error", err)
		return
	}
}
