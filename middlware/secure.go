package middlware

import (
	"silvex/app"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func secureApp(server *fiber.App) {

	// Attach limiter middleware
	if app.Config.Limiter.Enable {
		server.Use(getServerLimiter())
	}
}

func getServerLimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        app.Config.Limiter.Retries,
		Expiration: time.Duration(app.Config.Limiter.JailTime) * time.Second,
		LimitReached: func(ctx *fiber.Ctx) error {
			app.Logs.Info().Str("banned-client-ip", ctx.IP())
			return ctx.SendStatus(fiber.StatusTooManyRequests)
		},
	})
}
