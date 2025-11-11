package main

import (
	"embed"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/henyxia/satisfactory-megafactory/server"
)

//go:embed web
var webFS embed.FS

func main() {
	// setup logger
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Logger()
	logger.Info().Msg("Satisfactory Megafactory starting")

	srv, err := server.New(&logger, &webFS)
	if err != nil {
		logger.Fatal().Err(err).Msg("unable to create server")
	}

	srv.Run()
}
