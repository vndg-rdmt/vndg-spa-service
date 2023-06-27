package main

import (
	"silvex/app"
	"silvex/routing"

	"github.com/gofiber/fiber/v2"
)

type Webserver struct {
	engine                *fiber.App
	engineRouter          func(*fiber.App)
	startupFunc           func(*fiber.App) error
	callbackFailedToStart func(error)
}

func (s *Webserver) Launch() {
	s.engineRouter(s.engine)
	err := s.startupFunc(s.engine)
	if s.callbackFailedToStart != nil && err != nil {
		s.callbackFailedToStart(err)
	}
}

func main() {
	server := Webserver{
		engine:                EngineMultiprocessOptimizied(handlerLimiterBan(app.LogClientBannedByLimiter)),
		startupFunc:           StartupProduction,
		engineRouter:          routing.AttachClientRouting,
		callbackFailedToStart: CallbackFatalFailStartup(app.LogServerFailedToStart, eventDetailerFatal),
	}
	server.Launch()
}
