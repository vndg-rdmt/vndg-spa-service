package routing

import (
	"github.com/gofiber/fiber/v2"
)

func DefineRouting(server *fiber.App) error {
	err := serveClientApp(server)
	if err != nil {
		return err
	}

	return nil
}
