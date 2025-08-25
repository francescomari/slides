package out_test

import (
	"log/slog"
	"testing"
)

// BEGIN OMIT
func TestOutput(t *testing.T) {
	log := slog.New(slog.NewJSONHandler(t.Output(), nil))
	log.Info("hello, there!", slog.String("foo", "bar"))
}

// END OMIT
