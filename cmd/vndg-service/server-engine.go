package main

import (
	"silvex/pkg/app"
	"time"

	"strings"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func EngineMultiprocessOptimizied(limiterBanHandler fiber.Handler) *fiber.App {

	engine := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})

	engine.Use(limiter.New(limiter.Config{
		Max:          app.Config.Limiter.Retries,
		Expiration:   time.Duration(app.Config.Limiter.JailTime) * time.Second,
		LimitReached: limiterBanHandler,
	}))

	return engine
}

func StartupProduction(engine *fiber.App) error {
	return engine.Listen(strings.Join([]string{app.Config.Server.Addr, app.Config.Server.Port}, ":"))
}
