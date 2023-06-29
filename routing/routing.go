package routing

import (
	"silvex/pkg/app"

	"github.com/gofiber/fiber/v2"
)

func SinglePageApplication(appHandler fiber.Handler) func(server *fiber.App) {
	return func(server *fiber.App) {
		server.Static("/", app.Config.Client.Directory)
		server.Get("/*", appHandler)
	}
}
