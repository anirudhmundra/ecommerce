package main

import (
	"bcg/ecommerce/config"
	"bcg/ecommerce/data"
	"os"

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
}
