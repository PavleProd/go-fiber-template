package handlers

import (
	"github.com/PavleProd/go-fiber-template/pkg/container"
	"github.com/PavleProd/go-fiber-template/pkg/handlers/example"
	"github.com/gofiber/fiber/v3"
)

func Routes(root fiber.Router, cont *container.Container) {
	example.Routes(root, cont)
}
