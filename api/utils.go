package api

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/cleysonph/users-api/db"
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

func sendError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	if errors.Is(err, errEmptyFirstName) || errors.Is(err, errEmptyLastName) || errors.Is(err, errEmptyBiography) || errors.Is(err, db.ErrInvalidUUID) {
		sendJSON(w, response{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	if errors.Is(err, db.ErrUserNotFound) {
		sendJSON(w, response{Error: err.Error()}, http.StatusNotFound)
		return
	}

	sendJSON(w, response{Error: err.Error()}, http.StatusInternalServerError)
}
