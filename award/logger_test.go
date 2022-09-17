package award_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/haleyrc/cheevos/award"
	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestLoggerLogsAnErrorFromAwardCheevoToUser(t *testing.T) {
	logger := testutil.NewTestLogger()

	al := &award.Logger{
		Svc: &mock.AwardService{
			AwardCheevoToUserFn: func(_ context.Context, _, _ string) error { return fmt.Errorf("oops") },
		},
		Logger: logger,
	}
	al.AwardCheevoToUser(context.Background(), "userid", "cheevoid")

	logger.ShouldLog(t,
		`{"Fields":{"Cheevo":"cheevoid","User":"userid"},"Message":"awarding cheevo to user"}`,
		`{"Fields":{"Error":"oops"},"Message":"award cheevo to user failed"}`,
	)
}

func TestLoggerLogsTheResponseFromFromAwardCheevoToUser(t *testing.T) {
	logger := testutil.NewTestLogger()

	al := &award.Logger{
		Svc: &mock.AwardService{
			AwardCheevoToUserFn: func(_ context.Context, _, _ string) error { return nil },
		},
		Logger: logger,
	}
	al.AwardCheevoToUser(context.Background(), "userid", "cheevoid")

	logger.ShouldLog(t,
		`{"Fields":{"Cheevo":"cheevoid","User":"userid"},"Message":"awarding cheevo to user"}`,
		`{"Fields":{"Cheevo":"cheevoid","User":"userid"},"Message":"awarded cheevo to user"}`,
	)
}
