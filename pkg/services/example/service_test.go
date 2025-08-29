package example_test

import (
	"github.com/catadoo/daily-mini-games/pkg/services/example"
	"github.com/catadoo/daily-mini-games/pkg/testutils"
	"github.com/rs/zerolog"
	"testing"
)

func TestLog(t *testing.T) {
	t.Parallel()

	logger, tester := testutils.NewAppTestLogger(t, zerolog.InfoLevel)
	service := example.New(logger)

	service.Log("hi")

	tester.Entries().ExpLen(1)
	tester.FirstEntry().ExpMsg("example service called with message hi")
}
