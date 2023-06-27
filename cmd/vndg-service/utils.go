package main

import (
	"os"
	"silvex/app"
	"syscall"

	"github.com/gofiber/fiber/v2"
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

func handlerLimiterBan(eventName string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		app.Logs.Info().Str(eventName, ctx.IP()).Send()
		return ctx.SendStatus(fiber.StatusTooManyRequests)
	}
}
