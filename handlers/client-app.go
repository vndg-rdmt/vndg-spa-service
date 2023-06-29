package handlers

import (
	"fmt"
	"os"
	"path"
	"silvex/pkg/app"

	"github.com/gofiber/fiber/v2"
)

func ClientAppHandler() fiber.Handler {
	index := path.Join(app.Config.Client.Directory, app.Config.Client.Document)
	return func(c *fiber.Ctx) error {
		return c.SendFile(index)
	}
}

func ClientAppHandlerCached() fiber.Handler {
	if buffer, err := os.ReadFile(path.Join(app.Config.Client.Directory, app.Config.Client.Document)); err != nil {
		panic(fmt.Sprintf("Cannot load client app document,\nerror - %e", err))
	} else {
		return func(c *fiber.Ctx) error {
			return c.Send(buffer)
		}
	}

}
