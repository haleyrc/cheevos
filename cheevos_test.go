package cheevos_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/haleyrc/cheevos/log"
)

func NewTestLogger() *TestLogger {
	var buff bytes.Buffer
	logger := &log.JSONLogger{
		EnableDebug: true,
		Output:      &buff,
	}
	if testing.Verbose() {
		logger.Output = io.MultiWriter(&buff, os.Stdout)
	}

	return &TestLogger{
		JSONLogger: logger,
		buff:       &buff,
	}
}

type TestLogger struct {
	*log.JSONLogger
	buff *bytes.Buffer
}

func (tl *TestLogger) ShouldLog(t *testing.T, want ...string) {
	t.Helper()
	diff(t, join(want...), strings.TrimSpace(tl.buff.String()))
}

func diff(t *testing.T, want, got string) bool {
	t.Helper()
	if want != got {
		t.Errorf("logger output mismatch\nwant:\n%s\ngot:\n%s", want, got)
		return false
	}
	return true
}

func join(ss ...string) string {
	return strings.Join(ss, "\n")
}
