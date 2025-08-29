package testutils

import (
	"github.com/rs/zerolog"
	"github.com/rzajac/zltest"
	"testing"
)

func NewAppTestLogger(tb testing.TB, level zerolog.Level) (zerolog.Logger, *zltest.Tester) {
	tb.Helper()
	
	tst := zltest.New(tb)
	testWriter := zerolog.NewTestWriter(tb)
	logger := zerolog.New(zerolog.MultiLevelWriter(tst, testWriter)).Level(level)

	return logger, tst
}
