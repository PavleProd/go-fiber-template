package handlers

import (
	"github.com/catadoo/daily-mini-games/pkg/container"
	"github.com/catadoo/daily-mini-games/pkg/handlers/example"
	"github.com/gofiber/fiber/v3"
)

func Routes(root fiber.Router, cont *container.Container) {
	example.Routes(root, cont)
}
