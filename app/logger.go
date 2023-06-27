package app

import (
	"os"

	"github.com/rs/zerolog"
)

// Preallocated logger
var Logs zerolog.Logger

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	ConfigureDebugLogger()
}

func ConfigureDebugLogger() {
	Logs = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
}
