package invitation_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/invitation"
	"github.com/haleyrc/cheevos/lib/time"
)

/*
AcceptInvitation(ctx context.Context, userID, code string) error
DeclineInvitation(ctx context.Context, code string) error
InviteUserToOrganization(ctx context.Context, email, orgID string) (*Invitation, error)
RefreshInvitation(ctx context.Context, code string) error
*/

func TestLoggerLogsAnErrorFromAcceptInvitation(t *testing.T) {
	logger := testutil.NewTestLogger()

	il := &invitation.Logger{
		Svc: &mock.InvitationService{
			AcceptInvitationFn: func(_ context.Context, _, _ string) error { return fmt.Errorf("oops") },
		},
		Logger: logger,
	}
	il.AcceptInvitation(context.Background(), "userid", "code")

	logger.ShouldLog(t,
		`{"Fields":{"Code":"code","User":"userid"},"Message":"accepting invitation"}`,
		`{"Fields":{"Error":"oops"},"Message":"accept invitation failed"}`,
	)
}

func TestLoggerLogsTheResponseFromAcceptInvitation(t *testing.T) {
	logger := testutil.NewTestLogger()

	il := &invitation.Logger{
		Svc: &mock.InvitationService{
			AcceptInvitationFn: func(_ context.Context, _, _ string) error { return nil },
		},
		Logger: logger,
	}
	il.AcceptInvitation(context.Background(), "userid", "code")

	logger.ShouldLog(t,
		`{"Fields":{"Code":"code","User":"userid"},"Message":"accepting invitation"}`,
		`{"Fields":{"Code":"code","User":"userid"},"Message":"accepted invitation"}`,
	)
}

func TestLoggerLogsAnErrorFromDeclineInvitation(t *testing.T) {
	logger := testutil.NewTestLogger()

	il := &invitation.Logger{
		Svc: &mock.InvitationService{
			DeclineInvitationFn: func(_ context.Context, _ string) error { return fmt.Errorf("oops") },
		},
		Logger: logger,
	}
	il.DeclineInvitation(context.Background(), "code")

	logger.ShouldLog(t,
		`{"Fields":{"Code":"code"},"Message":"declining invitation"}`,
		`{"Fields":{"Error":"oops"},"Message":"decline invitation failed"}`,
	)
}

func TestLoggerLogsTheResponseFromDeclineInvitation(t *testing.T) {
	logger := testutil.NewTestLogger()

	il := &invitation.Logger{
		Svc: &mock.InvitationService{
			DeclineInvitationFn: func(_ context.Context, _ string) error { return nil },
		},
		Logger: logger,
	}
	il.DeclineInvitation(context.Background(), "code")

	logger.ShouldLog(t,
		`{"Fields":{"Code":"code"},"Message":"declining invitation"}`,
		`{"Fields":{"Code":"code"},"Message":"declined invitation"}`,
	)
}

func TestLoggerLogsAnErrorFromInviteUserToOrganization(t *testing.T) {
	logger := testutil.NewTestLogger()

	il := &invitation.Logger{
		Svc: &mock.InvitationService{
			InviteUserToOrganizationFn: func(_ context.Context, _, _ string) (*invitation.Invitation, error) { return nil, fmt.Errorf("oops") },
		},
		Logger: logger,
	}
	il.InviteUserToOrganization(context.Background(), "email", "orgid")

	logger.ShouldLog(t,
		`{"Fields":{"Email":"email","Organization":"orgid"},"Message":"inviting user to organization"}`,
		`{"Fields":{"Error":"oops"},"Message":"invite user to organization failed"}`,
	)
}

func TestLoggerLogsTheResponseFromInviteUserToOrganization(t *testing.T) {
	time.Freeze()

	ctx := context.Background()
	logger := testutil.NewTestLogger()
	inv := &invitation.Invitation{
		Email:          "email",
		OrganizationID: "orgid",
		Expires:        time.Now(),
	}

	il := &invitation.Logger{
		Svc: &mock.InvitationService{
			InviteUserToOrganizationFn: func(_ context.Context, _, _ string) (*invitation.Invitation, error) { return inv, nil },
		},
		Logger: logger,
	}
	il.InviteUserToOrganization(ctx, inv.Email, inv.OrganizationID)

	logger.ShouldLog(t,
		`{"Fields":{"Email":"email","Organization":"orgid"},"Message":"inviting user to organization"}`,
		`{"Fields":{"Invitation":{"Email":"email","OrganizationID":"orgid","Expires":"2022-09-16T15:02:04Z"}},"Message":"invited user to organization"}`,
	)
}

func TestLoggerLogsAnErrorFromRefreshInvitation(t *testing.T) {
	time.Freeze()

	logger := testutil.NewTestLogger()

	il := &invitation.Logger{
		Svc: &mock.InvitationService{
			RefreshInvitationFn: func(_ context.Context, _ string) error { return fmt.Errorf("oops") },
		},
		Logger: logger,
	}
	il.RefreshInvitation(context.Background(), "code")

	logger.ShouldLog(t,
		`{"Fields":{"Code":"code"},"Message":"refreshing invitation"}`,
		`{"Fields":{"Error":"oops"},"Message":"refresh invitation failed"}`,
	)
}

func TestLoggerLogsTheResponseFromRefreshInvitation(t *testing.T) {
	time.Freeze()

	logger := testutil.NewTestLogger()

	il := &invitation.Logger{
		Svc: &mock.InvitationService{
			RefreshInvitationFn: func(_ context.Context, _ string) error { return nil },
		},
		Logger: logger,
	}
	il.RefreshInvitation(context.Background(), "code")

	logger.ShouldLog(t,
		`{"Fields":{"Code":"code"},"Message":"refreshing invitation"}`,
		`{"Fields":{"Code":"code"},"Message":"refreshed invitation"}`,
	)
}
