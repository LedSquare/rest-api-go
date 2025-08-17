package main

import (
	"log/slog"
	"os"
	"rest-api-go/internal/config"
	"rest-api-go/internal/lib/logger/sl"
	"rest-api-go/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	config := config.MustLoad()

	log := setupLogger(config.Env)

	log.Info("Starting app", slog.String("env", config.Env))
	log.Debug("Debug mod is on")

	storage, err := sqlite.New(config.StoragePath)
	if err != nil {
		log.Error("Failed init storage", sl.Error(err))
		os.Exit(1)
	}

	_ = storage

	os.Exit(0)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(
			os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug},
		))
	case envDev:
		log = slog.New(slog.NewTextHandler(
			os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug},
		))
	case envProd:
		log = slog.New(slog.NewJSONHandler(
			os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo},
		))
	}

	return log
}
