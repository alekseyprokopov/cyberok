package main

import (
	"context"
	"cyberok/internal/config"
	"cyberok/internal/http-server/handlers/getFqdn"
	"cyberok/internal/http-server/handlers/getIp"
	"cyberok/internal/http-server/handlers/getWhois"
	"cyberok/internal/http-server/handlers/setFqdn"
	"cyberok/internal/http-server/handlers/setWhois"
	"cyberok/internal/lib/api/logger/sl"
	"cyberok/internal/repository"
	pg "cyberok/internal/repository/postgres"
	"cyberok/internal/resolvers/fqdn"
	"cyberok/internal/resolvers/whois"
	"cyberok/internal/service"
	"cyberok/internal/worker"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"golang.org/x/exp/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	db, err := pg.NewPostgresDB(cfg.DB)
	if err != nil {
		log.Error("cant init repository: ", err)
		os.Exit(1)
	}
	repository := repository.NewRepository(db)

	//migrations
	runMigration(cfg, log)

	//resolvers
	whoisResolver := whois.NewWhoisResolver(cfg.DNS)
	fqdnResolver, err := fqdn.NewDnsxResolver(cfg.DNS)
	if err != nil {
		log.Error("cant init fqdnResolver: ", err)
		os.Exit(1)
	}

	//services
	services := service.New(repository, fqdnResolver, whoisResolver)
	updateWorker := worker.NewUpdateWorker(cfg.DB, log, services)
	updateWorker.Start()

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

	log.Info("starting server", slog.String("port:", cfg.HTTPServer.Port))

	srv := &http.Server{
		Addr:         ":" + cfg.HTTPServer.Port,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err = srv.ListenAndServe(); err != nil {
			log.Error("error occurred on server shutting down: ", err)
		}
	}()

	log.Info("server started")

	//Shutdown
	<-done
	log.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.HTTPServer.StopTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", sl.Err(err))
		return
	}

	updateWorker.Close()

	if err := db.Close(); err != nil {
		log.Error("error occurred on db connection close:", sl.Err(err))
	}

	log.Info("server stopped")

}

func runMigration(cfg *config.Config, log *slog.Logger) {
	m, err := migrate.New(
		cfg.MigrationUrl,
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName, cfg.DB.SSLMode))

	if err != nil {
		log.Error("cannot create new migrate instance: ", err)
	}
	if err := m.Up(); err != nil {
		log.Error("failed to run migrate up: ", err)
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
