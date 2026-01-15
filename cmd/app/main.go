package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/tmozzze/ChatUs/internal/config"
	"github.com/tmozzze/ChatUs/internal/storage"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// Init Config
	cfg := config.MustLoad()

	// Init Logger
	log := setupLogger(envDev)
	log.Info("starting SkoobyTODO", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	// Init Storage
	_, err := storage.InitDB(cfg.Postgres)
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to init storage: %v\n", err)
		os.Exit(1)
	}
	log.Info("storage is initialized")

	// Init Repository

	// Init Services

	// Init Handlers

	// Strart Server
}

func setupLogger(env string) *slog.Logger {
	switch env {
	case envLocal: // Text Debug
		return slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev: // JSON Debug
		return slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd: // JSON Info
		return slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return slog.New(slog.NewTextHandler(os.Stdout, nil))
}
