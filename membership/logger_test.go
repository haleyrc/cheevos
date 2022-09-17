package membership_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/membership"
)

func TestLoggerLogsAnErrorFromAddMemberToOrganization(t *testing.T) {
	logger := testutil.NewTestLogger()

	ml := &membership.Logger{
		Svc: &mock.MembershipService{
			AddMemberToOrganizationFn: func(ctx context.Context, userID, orgID string) error { return fmt.Errorf("oops") },
		},
		Logger: logger,
	}
	ml.AddMemberToOrganization(context.Background(), "userid", "orgid")

	logger.ShouldLog(t,
		`{"Fields":{"Organization":"orgid","User":"userid"},"Message":"adding member to organization"}`,
		`{"Fields":{"Error":"oops"},"Message":"add member to organization failed"}`,
	)
}

func TestLoggerLogsTheResponseFromAddMemberToOrganization(t *testing.T) {
	logger := testutil.NewTestLogger()

	ml := &membership.Logger{
		Svc: &mock.MembershipService{
			AddMemberToOrganizationFn: func(ctx context.Context, userID, orgID string) error { return nil },
		},
		Logger: logger,
	}
	ml.AddMemberToOrganization(context.Background(), "userid", "orgid")

	logger.ShouldLog(t,
		`{"Fields":{"Organization":"orgid","User":"userid"},"Message":"adding member to organization"}`,
		`{"Fields":{"Organization":"orgid","User":"userid"},"Message":"added member to organization"}`,
	)
}
