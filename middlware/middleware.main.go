package middlware

import (
	"silvex/app"

	"github.com/gofiber/fiber/v2"
)

func AttachMiddlware(server *fiber.App, debug bool) {
	if debug {
		server.Use(debuggingMiddlware)
	}
	secureApp(server)
}

func debuggingMiddlware(ctx *fiber.Ctx) error {
	req := ctx.Request()
	app.Logs.Debug().
		Str("uri", string(req.URI().FullURI())).
		Str("headers", string(req.Header.Header())).
		Str("body", string(req.Body())).
		Str("new connection", ctx.IP()).
		Send()
	return ctx.Next()
}
