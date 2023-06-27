package app

import (
	"os"

	"github.com/rs/zerolog"
)

var Logs *zerolog.Logger
var _localV = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
var _debugV = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func ConfigureLocalLogger() {
	Logs = &_localV
}

func ConfigureDebugLogger() {
	Logs = &_debugV
}
