package logs

import (
	"os"

	"github.com/rs/zerolog"
)

// Preallocated logger
var Writer zerolog.Logger

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	SetupDebugLogger()
}

func SetupDebugLogger() {
	Writer = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
}
