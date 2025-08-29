package services

import (
	"github.com/catadoo/daily-mini-games/pkg/config"
	"github.com/catadoo/daily-mini-games/pkg/services/example"
	"github.com/rs/zerolog"
	"sync"
)

type (
	Services struct {
		Sync
		
		cfg    *config.Config
		logger zerolog.Logger

		example example.Interface
	}

	Sync struct {
		exampleServiceOnce sync.Once
	}
)

func New(cfg *config.Config, logger zerolog.Logger) *Services {
	return &Services{
		cfg:    cfg,
		logger: logger,
	}
}
