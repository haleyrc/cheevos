package cheevo_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/haleyrc/cheevos/cheevo"
	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestLoggerLogsAnErrorFromCreateCheevo(t *testing.T) {
	logger := testutil.NewTestLogger()

	cl := &cheevo.Logger{
		Svc: &mock.CheevoService{
			CreateCheevoFn: func(_ context.Context, _, _, _ string) (*cheevo.Cheevo, error) {
				return nil, fmt.Errorf("oops")
			},
		},
		Logger: logger,
	}
	cl.CreateCheevo(context.Background(), "name", "description", "orgid")

	logger.ShouldLog(t,
		`{"Fields":{"Description":"description","Name":"name","Organization":"orgid"},"Message":"creating cheevo"}`,
		`{"Fields":{"Error":"oops"},"Message":"create cheevo failed"}`,
	)
}

func TestLoggerLogsTheResponseFromCreateCheevo(t *testing.T) {
	logger := testutil.NewTestLogger()

	cl := &cheevo.Logger{
		Svc: &mock.CheevoService{
			CreateCheevoFn: func(_ context.Context, _, _, _ string) (*cheevo.Cheevo, error) {
				return &cheevo.Cheevo{ID: "id", Name: "name", Description: "description"}, nil
			},
		},
		Logger: logger,
	}
	cl.CreateCheevo(context.Background(), "name", "description", "orgid")

	logger.ShouldLog(t,
		`{"Fields":{"Description":"description","Name":"name","Organization":"orgid"},"Message":"creating cheevo"}`,
		`{"Fields":{"Cheevo":{"ID":"id","Name":"name","Description":"description"}},"Message":"cheevo created"}`,
	)
}
