package example_test

import (
	"github.com/PavleProd/go-fiber-template/pkg/services/example"
	"github.com/PavleProd/go-fiber-template/pkg/testutils"
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
