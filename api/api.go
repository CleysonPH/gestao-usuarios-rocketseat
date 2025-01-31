package api

import (
	"net/http"

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

	return r
}

func handleFindAllUsers(ur db.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := ur.FindAll()
		sendJSON(w, response{Data: users}, http.StatusOK)
	}
}
