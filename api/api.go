package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/cleysonph/users-api/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(ur db.UserRepository) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Get("/api/users", handleFindAllUsers(ur))
	r.Post("/api/users", handleInsertUser(ur))

	return r
}

func handleInsertUser(ur db.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u db.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			sendJSON(w, response{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		if len(strings.TrimSpace(u.FirstName)) <= 0 {
			sendJSON(w, response{Error: "please provide a first_name"}, http.StatusBadRequest)
			return
		}
		if len(strings.TrimSpace(u.LastName)) <= 0 {
			sendJSON(w, response{Error: "please provide a last_name"}, http.StatusBadRequest)
			return
		}
		if len(strings.TrimSpace(u.Biography)) <= 0 {
			sendJSON(w, response{Error: "please provide a biography"}, http.StatusBadRequest)
			return
		}

		user := ur.Insert(u)
		sendJSON(w, response{Data: user}, http.StatusCreated)
	}
}

func handleFindAllUsers(ur db.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := ur.FindAll()
		sendJSON(w, response{Data: users}, http.StatusOK)
	}
}
