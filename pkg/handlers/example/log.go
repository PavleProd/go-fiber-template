package example

import (
	"github.com/PavleProd/go-fiber-template/pkg/dto"
	exampleDto "github.com/PavleProd/go-fiber-template/pkg/dto/example"
	exampleService "github.com/PavleProd/go-fiber-template/pkg/services/example"
	"github.com/gofiber/fiber/v3"
)

func Log(service exampleService.Interface) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var request exampleDto.LogRequest

		if err := ctx.Bind().Body(&request); err != nil {
			return err
		}

		if err := request.Validate(); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorMessage{
				Message: err.Error(),
			})
		}

		service.Log(request.Message)

		return ctx.Status(fiber.StatusOK).End()
	}
}
