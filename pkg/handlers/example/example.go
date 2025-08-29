package example

import (
	"github.com/PavleProd/go-fiber-template/pkg/container"
	"github.com/gofiber/fiber/v3"
)

const Base = "/example"

func Routes(root fiber.Router, cont *container.Container) fiber.Router {
	example := root.Group(Base)

	example.Get("", Log(cont.GetExampleService()))

	return example
}
