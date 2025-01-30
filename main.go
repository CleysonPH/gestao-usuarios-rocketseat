package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/cleysonph/users-api/api"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}

	slog.Info("all systems offline")
}

func run() error {
	handler := api.NewHandler()

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
