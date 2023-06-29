package main

import (
	"os"
	"silvex/pkg/app"
	"silvex/pkg/logs"
	"syscall"

	"github.com/rs/zerolog"
)

func eventDetailerFatal(err error) *zerolog.Event {
	return zerolog.Dict().
		Int("user-id", syscall.Getuid()).
		Int("pid", os.Getegid()).
		Str("addr", app.Config.Server.Addr).
		Str("port", app.Config.Server.Port).
		Err(err)
}

func callbackFatalFailStartup(eventName string, eventDetailer func(error) *zerolog.Event) func(error) {
	return func(err error) {
		logs.Writer.Fatal().
			Str(logs.MarkerEvent, eventName).
			Dict(logs.MarkerDetails, eventDetailer(err)).
			Send()
	}
}
