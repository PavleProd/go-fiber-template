package core

import (
	"context"
	"fmt"
	"github.com/catadoo/daily-mini-games/pkg/container"
	"github.com/catadoo/daily-mini-games/pkg/handlers"
	"github.com/gofiber/fiber/v3"
)

type Core struct {
	app  *fiber.App
	cont *container.Container
}

func New(cont *container.Container) *Core {
	errorHandlerLogger := cont.GetLogger().With().Str("handler", "error-handler").Logger()
	fiberCfg := fiber.Config{
		ErrorHandler: handlers.ErrorHandler(errorHandlerLogger),
	} // TODO: config fiber

	return &Core{
		app:  fiber.New(fiberCfg),
		cont: cont,
	}
}

func (c *Core) RegisterRoutes() {
	handlers.Routes(c.app, c.cont)
}

func (c *Core) Listen() error {
	cfg := c.cont.GetConfig()

	address := fmt.Sprintf("%s:%d", cfg.Application.Host, cfg.Application.Port)

	if err := c.app.Listen(address); err != nil {
		return err
	}

	return nil
}

func (c *Core) Close(ctx context.Context) error {
	if err := c.app.ShutdownWithContext(ctx); err != nil {
		return err
	}

	return nil
}
