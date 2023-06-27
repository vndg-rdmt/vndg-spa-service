package routing

import (
	"github.com/gofiber/fiber/v2"
)

func AttachClientRouting(server *fiber.App) {
	// server.Static("/", app.Config.Client.Directory)
	server.Get("/*", handleClientApp)
}

func handleClientApp(ctx *fiber.Ctx) error {
	return ctx.Send(ctx.Request().Header.Header())
}
