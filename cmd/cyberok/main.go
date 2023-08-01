package main

import (
	"cyberok/internal/config"
	"fmt"
	"golang.org/x/exp/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
)

func main() {
	//config
	cfg := config.MustLoad()
	fmt.Println(cfg)

	//logger
	//log := setupLogger(envLocal)
	//log.Info("starting wtf")
	//log.Debug("debug messages")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}
	return log
}
