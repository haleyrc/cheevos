package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/haleyrc/cheevos/cheevos"
	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestLoggerLogsAnErrorFromAwardCheevoToUser(t *testing.T) {
	logger := testutil.NewTestLogger()

	al := &cheevos.Logger{
		Service: &mock.CheevosService{
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

	al := &cheevos.Logger{
		Service: &mock.CheevosService{
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

func TestLoggerLogsAnErrorFromCreateCheevo(t *testing.T) {
	logger := testutil.NewTestLogger()

	cl := &cheevos.Logger{
		Service: &mock.CheevosService{
			CreateCheevoFn: func(_ context.Context, _, _, _ string) (*cheevos.Cheevo, error) {
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

	cl := &cheevos.Logger{
		Service: &mock.CheevosService{
			CreateCheevoFn: func(_ context.Context, _, _, _ string) (*cheevos.Cheevo, error) {
				return &cheevos.Cheevo{ID: "id", Name: "name", Description: "description", OrganizationID: "orgid"}, nil
			},
		},
		Logger: logger,
	}
	cl.CreateCheevo(context.Background(), "name", "description", "orgid")

	logger.ShouldLog(t,
		`{"Fields":{"Description":"description","Name":"name","Organization":"orgid"},"Message":"creating cheevo"}`,
		`{"Fields":{"Cheevo":{"ID":"id","Name":"name","Description":"description","OrganizationID":"orgid"}},"Message":"cheevo created"}`,
	)
}
