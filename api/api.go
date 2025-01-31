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
	r.Get("/api/users/{id}", handleFindUserById(ur))
	r.Delete("/api/users/{id}", handleDeleteUserById(ur))
	r.Put("/api/users/{id}", handleUpdateUserById(ur))

	return r
}

func handleUpdateUserById(ur db.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		var u db.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			sendJSON(w, response{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		err := validateUserData(u)
		if err != nil {
			sendError(w, err)
			return
		}

		user, err := ur.UpdateById(id, u)
		if err != nil {
			sendError(w, err)
			return
		}

		sendJSON(w, response{Data: user}, http.StatusCreated)
	}
}

func handleDeleteUserById(ur db.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if err := ur.DeleteById(id); err != nil {
			sendError(w, err)
			return
		}

		sendJSON(w, response{}, http.StatusNoContent)
	}
}

func handleFindUserById(ur db.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		user, err := ur.FindById(id)
		if err != nil {
			sendError(w, err)
			return
		}

		sendJSON(w, response{Data: user}, http.StatusOK)
	}
}

func handleInsertUser(ur db.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u db.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			sendJSON(w, response{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		err := validateUserData(u)
		if err != nil {
			sendError(w, err)
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

func validateUserData(u db.User) error {
	if len(strings.TrimSpace(u.FirstName)) <= 0 {
		return errEmptyFirstName
	}
	if len(strings.TrimSpace(u.LastName)) <= 0 {
		return errEmptyLastName
	}
	if len(strings.TrimSpace(u.Biography)) <= 0 {
		return errEmptyBiography
	}
	return nil
}
