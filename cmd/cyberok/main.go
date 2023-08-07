package main

import (
	"cyberok/internal/config"
	"cyberok/internal/http-server/handlers/getFqdn"
	"cyberok/internal/http-server/handlers/getIp"
	"cyberok/internal/http-server/handlers/getWhois"
	"cyberok/internal/http-server/handlers/setFqdn"
	"cyberok/internal/http-server/handlers/setWhois"
	"cyberok/internal/repository"
	"cyberok/internal/repository/postgres"
	"cyberok/internal/resolvers/fqdn"
	"cyberok/internal/resolvers/whois"
	"cyberok/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"golang.org/x/exp/slog"
	"net/http"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
)

func main() {
	//config
	cfg := config.MustLoad()

	//logger
	log := setupLogger(envLocal)

	//repository
	db, err := postgres.NewPostgresDB(cfg.DB)
	if err != nil {
		log.Error("cant init repository: ", err)
		os.Exit(1)
	}
	repository := repository.NewRepository(db)

	//resolvers
	whoisResolver := whois.NewWhoisResolver(cfg.DNS)
	fqdnResolver, err := fqdn.NewDnsxResolver(cfg.DNS)
	if err != nil {
		log.Error("cant init fqdnResolver: ", err)
		os.Exit(1)
	}

	//services
	services := service.New(repository, fqdnResolver, whoisResolver)

	//router
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/fqdn", getFqdn.New(log, services))
	router.Post("/ip", getIp.New(log, services))
	router.Post("/whois", getWhois.New(log, services))

	router.Post("/admin/fqdn", setFqdn.New(log, services))
	router.Post("/admin/whois", setWhois.New(log, services))

	log.Info("starting server", slog.String("address", cfg.HTTPServer.Address))

	srv := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}
	if err = srv.ListenAndServe(); err != nil {
		log.Error("failed to start sever: ", err)
	}

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
