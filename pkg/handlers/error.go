package handlers

import (
	"errors"
	"github.com/PavleProd/go-fiber-template/pkg/constants"
	"github.com/PavleProd/go-fiber-template/pkg/dto"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
)

func ErrorHandler(logger zerolog.Logger) fiber.ErrorHandler {
	return func(ctx fiber.Ctx, err error) error {
		ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(dto.ErrorMessage{
				Message: fiberErr.Message,
			})
		}

		logger.Error().Err(err).Msg("unprocessable error returned")

		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorMessage{
			Message: constants.InternalServerError,
		})
	}
}
