package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/haleyrc/pkg/time"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestLoggerLogsAnErrorFromAcceptInvitation(t *testing.T) {
	logger := testutil.NewTestLogger()

	il := &rosterLogger{
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

	il := &rosterLogger{
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

	il := &rosterLogger{
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

	il := &rosterLogger{
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

	il := &rosterLogger{
		Service: &mock.RosterService{
			InviteUserToOrganizationFn: func(_ context.Context, _, _ string) (*cheevos.Invitation, error) { return nil, fmt.Errorf("oops") },
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
	inv := &cheevos.Invitation{
		ID:             "id",
		Email:          "email",
		OrganizationID: "orgid",
		Expires:        time.Now(),
	}

	il := &rosterLogger{
		Service: &mock.RosterService{
			InviteUserToOrganizationFn: func(_ context.Context, _, _ string) (*cheevos.Invitation, error) { return inv, nil },
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

	il := &rosterLogger{
		Service: &mock.RosterService{
			RefreshInvitationFn: func(_ context.Context, _ string) (*cheevos.Invitation, error) { return nil, fmt.Errorf("oops") },
		},
		Logger: logger,
	}
	il.RefreshInvitation(context.Background(), "id")

	logger.ShouldLog(t,
		`{"Fields":{"ID":"id"},"Message":"refreshing invitation"}`,
		`{"Fields":{"Error":"oops"},"Message":"refresh invitation failed"}`,
	)
}

func TestLoggerLogsTheResponseFromRefreshInvitation(t *testing.T) {
	logger := testutil.NewTestLogger()

	il := &rosterLogger{
		Service: &mock.RosterService{
			RefreshInvitationFn: func(_ context.Context, _ string) (*cheevos.Invitation, error) {
				return &cheevos.Invitation{ID: "id", Email: "email", OrganizationID: "orgid", Expires: time.Now()}, nil
			},
		},
		Logger: logger,
	}
	il.RefreshInvitation(context.Background(), "id")

	logger.ShouldLog(t,
		`{"Fields":{"ID":"id"},"Message":"refreshing invitation"}`,
		`{"Fields":{"Invitation":{"ID":"id","Email":"email","OrganizationID":"orgid","Expires":"2022-09-16T15:02:04Z"}},"Message":"refreshed invitation"}`,
	)
}

func TestLoggerLogsAnErrorFromCreateOrganization(t *testing.T) {
	logger := testutil.NewTestLogger()

	ol := &rosterLogger{
		Service: &mock.RosterService{
			CreateOrganizationFn: func(_ context.Context, name, ownerID string) (*cheevos.Organization, error) {
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

	ol := &rosterLogger{
		Service: &mock.RosterService{
			CreateOrganizationFn: func(_ context.Context, name, ownerID string) (*cheevos.Organization, error) {
				return &cheevos.Organization{ID: "id", Name: "name", OwnerID: "ownerid"}, nil
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
