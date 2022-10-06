package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/lib/time"
	"github.com/haleyrc/cheevos/roster"
)

func TestLoggerLogsAnErrorFromAcceptInvitation(t *testing.T) {
	logger := testutil.NewTestLogger()

	il := &roster.Logger{
		Service: &mock.RosterService{
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

	il := &roster.Logger{
		Service: &mock.RosterService{
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

	il := &roster.Logger{
		Service: &mock.RosterService{
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

	il := &roster.Logger{
		Service: &mock.RosterService{
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

	il := &roster.Logger{
		Service: &mock.RosterService{
			InviteUserToOrganizationFn: func(_ context.Context, _, _ string) (*roster.Invitation, error) { return nil, fmt.Errorf("oops") },
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
	ctx := context.Background()
	logger := testutil.NewTestLogger()
	inv := &roster.Invitation{
		ID:             "id",
		Email:          "email",
		OrganizationID: "orgid",
		Expires:        time.Now(),
	}

	il := &roster.Logger{
		Service: &mock.RosterService{
			InviteUserToOrganizationFn: func(_ context.Context, _, _ string) (*roster.Invitation, error) { return inv, nil },
		},
		Logger: logger,
	}
	il.InviteUserToOrganization(ctx, inv.Email, inv.OrganizationID)

	logger.ShouldLog(t,
		`{"Fields":{"Email":"email","Organization":"orgid"},"Message":"inviting user to organization"}`,
		`{"Fields":{"Invitation":{"ID":"id","Email":"email","OrganizationID":"orgid","Expires":"2022-09-16T15:02:04Z"}},"Message":"invited user to organization"}`,
	)
}

func TestLoggerLogsAnErrorFromRefreshInvitation(t *testing.T) {
	logger := testutil.NewTestLogger()

	il := &roster.Logger{
		Service: &mock.RosterService{
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
	logger := testutil.NewTestLogger()

	il := &roster.Logger{
		Service: &mock.RosterService{
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

func TestLoggerLogsAnErrorFromAddMemberToOrganization(t *testing.T) {
	logger := testutil.NewTestLogger()

	ml := &roster.Logger{
		Service: &mock.RosterService{
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

	ml := &roster.Logger{
		Service: &mock.RosterService{
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

func TestLoggerLogsAnErrorFromCreateOrganization(t *testing.T) {
	logger := testutil.NewTestLogger()

	ol := &roster.Logger{
		Service: &mock.RosterService{
			CreateOrganizationFn: func(_ context.Context, name, ownerID string) (*roster.Organization, error) {
				return nil, fmt.Errorf("oops")
			},
		},
		Logger: logger,
	}
	ol.CreateOrganization(context.Background(), "name", "ownerid")

	logger.ShouldLog(t,
		`{"Fields":{"Name":"name","Owner":"ownerid"},"Message":"creating organization"}`,
		`{"Fields":{"Error":"oops"},"Message":"create organization failed"}`,
	)
}

func TestLoggerLogsTheResponseFromCreateOrganization(t *testing.T) {
	logger := testutil.NewTestLogger()

	ol := &roster.Logger{
		Service: &mock.RosterService{
			CreateOrganizationFn: func(_ context.Context, name, ownerID string) (*roster.Organization, error) {
				return &roster.Organization{ID: "id", Name: "name", OwnerID: "ownerid"}, nil
			},
		},
		Logger: logger,
	}
	ol.CreateOrganization(context.Background(), "name", "ownerid")

	logger.ShouldLog(t,
		`{"Fields":{"Name":"name","Owner":"ownerid"},"Message":"creating organization"}`,
		`{"Fields":{"Organization":{"ID":"id","Name":"name","OwnerID":"ownerid"}},"Message":"organization created"}`,
	)
}
