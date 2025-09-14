package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"rest-api-go/internal/config"
	"rest-api-go/internal/http/handlers/url/save"
	middlewareLogger "rest-api-go/internal/http/middleware/logger"
	slogpretty "rest-api-go/internal/lib/logger/handlers"
	"rest-api-go/internal/lib/logger/sl"
	"rest-api-go/internal/storage/sqlite"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middlewareLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/url", save.New(log, storage))

	server := &http.Server{
		Addr:         config.Address,
		Handler:      router,
		ReadTimeout:  config.Timeout,
		WriteTimeout: config.Timeout,
		IdleTimeout:  config.IdleTimeout,
	}

	log.Info("Starting server", slog.String("address", config.Address))

	if err := server.ListenAndServe(); err != nil {
		log.Error("Failed to start server")
	}

	log.Error("Server is stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = setupPrettySlog()
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

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

func maxFreqSum(s string) int {
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	freq := make(map[rune]int)

	for _, ch := range s {
		freq[ch]++
	}

	fmt.Println(freq)
	maxVowelFreq := 0
	maxConsonantFreq := 0

	for ch, count := range freq {
		if vowels[ch] {
			if count > maxVowelFreq {
				maxVowelFreq = count
			}
		} else {
			if count > maxConsonantFreq {
				maxConsonantFreq = count
			}
		}
	}

	return maxVowelFreq + maxConsonantFreq
}
