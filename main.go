package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/cleysonph/users-api/api"
	"github.com/cleysonph/users-api/db"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}

	slog.Info("all systems offline")
}

func run() error {
	ur := db.NewUserRepository()

	handler := api.NewHandler(ur)

	srv := http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
