package main

import (
	"silvex/handlers"
	"silvex/pkg/logs"
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
	if err != nil && s.callbackFailedToStart != nil {
		s.callbackFailedToStart(err)
	}
}

func main() {
	server := Webserver{
		engine: EngineMultiprocessOptimizied(handlers.RejectClientDetailed(
			logs.EventClientRejected,
			fiber.StatusTooManyRequests,
		)),
		startupFunc:           StartupProduction,
		engineRouter:          routing.SinglePageApplication(handlers.ClientAppHandlerCached()),
		callbackFailedToStart: callbackFatalFailStartup(logs.EventServerFailedToStart, eventDetailerFatal),
	}
	server.Launch()
}
