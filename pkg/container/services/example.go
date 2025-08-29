package services

import "github.com/catadoo/daily-mini-games/pkg/services/example"

func (s *Services) GetExampleService() example.Interface {
	s.exampleServiceOnce.Do(func() {
		logger := s.logger.With().Str("service", "example").Logger()
		s.example = example.New(logger)
	})

	return s.example
}
