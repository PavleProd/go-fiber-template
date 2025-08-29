package container

import (
	"github.com/catadoo/daily-mini-games/pkg/config"
	"github.com/catadoo/daily-mini-games/pkg/container/services"
	"github.com/rs/zerolog"
)

type Container struct {
	*services.Services
	
	cfg    *config.Config
	logger zerolog.Logger
}

func New(cfg *config.Config) *Container {
	cont := Container{
		cfg:    cfg,
		logger: initLogger(),
	}

	cont.Services = services.New(cont.cfg, cont.logger)

	return &cont
}

func (c *Container) GetConfig() *config.Config {
	return c.cfg
}

func (c *Container) GetLogger() zerolog.Logger {
	return c.logger
}

func initLogger() zerolog.Logger {
	w := zerolog.NewConsoleWriter()

	return zerolog.New(w)
}
