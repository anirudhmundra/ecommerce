package main

import (
	"bcg/ecommerce/config"
	"bcg/ecommerce/data"
	entities "bcg/ecommerce/data/entities"
	"bcg/ecommerce/router"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout)
	logger.Info().Msg("ecommerce-bcg...")

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatal().Msgf("unable to load config :%-v", err)
	}
	items, err := data.LoadData(cfg.DataLoadFilePath)
	if err != nil {
		logger.Fatal().Msgf("unable to load data :%-v", err)
	}
	logger.Info().Msgf("%d items have been loaded", len(items))
	run(items, logger)
}

func run(items map[string]entities.ItemEntity, logger zerolog.Logger) {
	r := router.SetupRouter(items)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Info().Msgf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal().Msgf("Server forced to shutdown: %-v", err)
	}

	logger.Info().Msg("Server exiting")
}
