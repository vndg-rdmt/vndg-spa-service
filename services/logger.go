package service

import "github.com/rs/zerolog"

func ConfigureLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
