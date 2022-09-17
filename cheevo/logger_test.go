package cheevo_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/haleyrc/cheevos/cheevo"
	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestCheevoLoggerLogsAnErrorFromCreateCheevo(t *testing.T) {
	logger := testutil.NewTestLogger()

	cl := &cheevo.Logger{
		Svc: &mock.CheevoService{
			CreateCheevoFn: func(_ context.Context, _, _, _ string) (*cheevo.Cheevo, error) {
				return nil, fmt.Errorf("oops")
			},
		},
		Logger: logger,
	}
	cl.CreateCheevo(context.Background(), "Test", "This is a test.", "783bf2de-dce2-4f32-9f18-f77b904f87c")

	logger.ShouldLog(t,
		`{"Fields":{"Description":"This is a test.","Name":"Test","Organization":"783bf2de-dce2-4f32-9f18-f77b904f87c"},"Message":"creating cheevo"}`,
		`{"Fields":{"Error":"oops"},"Message":"create cheevo failed"}`,
	)
}

func TestCheevoLoggerLogsTheResponseFromCreateCheevo(t *testing.T) {
	logger := testutil.NewTestLogger()

	cl := &cheevo.Logger{
		Svc: &mock.CheevoService{
			CreateCheevoFn: func(_ context.Context, _, _, _ string) (*cheevo.Cheevo, error) {
				return &cheevo.Cheevo{
					ID:          "8059dcd7-bcc1-46fa-bfc0-3926c0b2c6ea",
					Name:        "Test",
					Description: "This is a test.",
				}, nil
			},
		},
		Logger: logger,
	}
	cl.CreateCheevo(context.Background(), "Test", "This is a test.", "238cb95f-8bcd-4cda-8cfc-9d03fecba894")

	logger.ShouldLog(t,
		`{"Fields":{"Description":"This is a test.","Name":"Test","Organization":"238cb95f-8bcd-4cda-8cfc-9d03fecba894"},"Message":"creating cheevo"}`,
		`{"Fields":{"Cheevo":{"ID":"8059dcd7-bcc1-46fa-bfc0-3926c0b2c6ea","Name":"Test","Description":"This is a test."}},"Message":"cheevo created"}`,
	)
}
