package main

import (
	"silvex/app"
	"strings"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type webserverStartupFunc func(*fiber.App) error

type Webserver struct {
	server      *fiber.App
	startupFunc webserverStartupFunc

	eventCallbackFailedToStart func(error)
}

func (s *Webserver) Launch() {
	err := s.startupFunc(s.server)
	if err != nil {
		return

	}
	return
}

func callbackFatalFail(eventName string) func(error) {
	return func(err error) {
		app.Logs.Fatal().Str(LogEvent, eventName).Dict().Send()
	}
}

/*
Create a new webserver.

All webserver configuration must be done withing this function,
configuration only includes things which affects webserver's engine,
and not it's behavior.
*/
func optimizedMultiprocessServer() *fiber.App {
	return fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})
}

/*
Launch server.

This function must contain all operations to be done after configuration
and functions, which changes webserver behavior or state, to make webserver
able to be up and running and make it active.

Webserver startup status is returned
*/
func productionStartup(server *fiber.App) error {
	err := server.Listen(strings.Join([]string{app.Config.Server.Addr, app.Config.Server.Port}, ":"))
	if err != nil {
		return err
	}
	return nil
}

/*
Server connections limiter.
Configuration fully based on app config
*/
func defaultLimiter() fiber.Handler {
	return limiter.New(limiter.Config{

		Max:        app.Config.Limiter.Retries,
		Expiration: time.Duration(app.Config.Limiter.JailTime) * time.Second,
		LimitReached: func(ctx *fiber.Ctx) error {

			app.Logs.Info().Str(LogClientBanned, ctx.IP())
			return ctx.SendStatus(fiber.StatusTooManyRequests)
		},
	})
}
