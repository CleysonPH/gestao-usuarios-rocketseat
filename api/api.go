package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewHandler() http.Handler {
	r := chi.NewMux()

	return r
}
