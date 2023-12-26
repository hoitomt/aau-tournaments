package main

import (
	"flag"
	"os"
	"path/filepath"

	exposureevents "github.com/aau-tournaments/exposure_events"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	debug := flag.Bool("debug", false, "set the log level to debug")

	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	exposureEventsData := filepath.Join("data", "exposure_events")
	err := os.MkdirAll(exposureEventsData, os.ModePerm)
	if err != nil {
		log.Fatal().Msgf("create exposure events data dir: %s", err)
	}
}

func main() {
	exposureevents.GetTournamentList()
}
