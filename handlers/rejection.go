package handlers

import (
	"silvex/pkg/logs"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

func RejectClientDetailed(byService string, rejectStatus int) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logs.Writer.Info().
			Str(logs.MarkerEvent, logs.EventClientRejected).
			Dict(logs.MarkerDetails, zerolog.Dict().
				Str(logs.FieldCaused, byService).
				Str(logs.FieldIp, ctx.IP())).
			Send()
		return ctx.SendStatus(rejectStatus)
	}
}
