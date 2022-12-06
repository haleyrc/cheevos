package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/haleyrc/cheevos/domain"
	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestLoggerLogsAnErrorFromAwardCheevoToUser(t *testing.T) {
	logger := testutil.NewTestLogger()

	al := &cheevosLogger{
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

	al := &cheevosLogger{
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

	cl := &cheevosLogger{
		Service: &mock.CheevosService{
			CreateCheevoFn: func(_ context.Context, _, _, _ string) (*domain.Cheevo, error) {
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

	cl := &cheevosLogger{
		Service: &mock.CheevosService{
			CreateCheevoFn: func(_ context.Context, _, _, _ string) (*domain.Cheevo, error) {
				return &domain.Cheevo{ID: "id", Name: "name", Description: "description", OrganizationID: "orgid"}, nil
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

func TestLoggerLogsAnErrorFromGetCheevo(t *testing.T) {
	logger := testutil.NewTestLogger()

	cl := cheevosLogger{
		Service: &mock.CheevosService{
			GetCheevoFn: func(_ context.Context, _ string) (*domain.Cheevo, error) {
				return nil, fmt.Errorf("oops")
			},
		},
		Logger: logger,
	}
	cl.GetCheevo(context.Background(), "id")

	logger.ShouldLog(t,
		`{"Fields":{"ID":"id"},"Message":"getting cheevo"}`,
		`{"Fields":{"Error":"oops"},"Message":"get cheevo failed"}`,
	)
}

func TestLoggerLogsTheReponseFromGetCheevo(t *testing.T) {
	logger := testutil.NewTestLogger()

	cl := cheevosLogger{
		Service: &mock.CheevosService{
			GetCheevoFn: func(_ context.Context, _ string) (*domain.Cheevo, error) {
				return &domain.Cheevo{ID: "id", Name: "name", Description: "description", OrganizationID: "orgid"}, nil
			},
		},
		Logger: logger,
	}
	cl.GetCheevo(context.Background(), "id")

	logger.ShouldLog(t,
		`{"Fields":{"ID":"id"},"Message":"getting cheevo"}`,
		`{"Fields":{"Cheevo":{"ID":"id","Name":"name","Description":"description","OrganizationID":"orgid"}},"Message":"got cheevo"}`,
	)
}
