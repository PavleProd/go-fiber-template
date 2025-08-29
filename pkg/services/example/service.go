package example

import "github.com/rs/zerolog"

type (
	Interface interface {
		Log(message string)
	}

	Service struct {
		logger zerolog.Logger
	}
)

var _ Interface = &Service{}

func New(logger zerolog.Logger) *Service {
	return &Service{
		logger: logger,
	}
}

func (s *Service) Log(message string) {
	s.logger.Info().Msgf("example service called with message %s", message)
}
