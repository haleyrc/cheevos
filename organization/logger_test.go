package organization_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/lib/time"
	"github.com/haleyrc/cheevos/organization"
)

func TestLoggerLogsAnErrorFromAddUserToOrganization(t *testing.T) {
	logger := testutil.NewTestLogger()

	ol := &organization.Logger{
		Svc: &mock.OrganizationService{
			AddMemberToOrganizationFn: func(_ context.Context, _, _ string) error { return fmt.Errorf("oops") },
		},
		Logger: logger,
	}
	ol.AddMemberToOrganization(context.Background(), "userid", "orgid")

	logger.ShouldLog(t,
		`{"Fields":{"Organization":"orgid","User":"userid"},"Message":"adding member to organization"}`,
		`{"Fields":{"Error":"oops"},"Message":"add member to organization failed"}`,
	)
}

func TestLoggerLogsTheResponseFromAddUserToOrganization(t *testing.T) {
	logger := testutil.NewTestLogger()

	ol := &organization.Logger{
		Svc: &mock.OrganizationService{
			AddMemberToOrganizationFn: func(_ context.Context, userID, ownerID string) error { return nil },
		},
		Logger: logger,
	}
	ol.AddMemberToOrganization(context.Background(), "userid", "orgid")

	logger.ShouldLog(t,
		`{"Fields":{"Organization":"orgid","User":"userid"},"Message":"adding member to organization"}`,
		`{"Fields":{"Organization":"orgid","User":"userid"},"Message":"added member to organization"}`,
	)
}

func TestLoggerLogsAnErrorFromCreateOrganization(t *testing.T) {
	logger := testutil.NewTestLogger()

	ol := &organization.Logger{
		Svc: &mock.OrganizationService{
			CreateOrganizationFn: func(_ context.Context, name, ownerID string) (*organization.Organization, error) {
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
	time.Freeze()

	logger := testutil.NewTestLogger()

	ol := &organization.Logger{
		Svc: &mock.OrganizationService{
			CreateOrganizationFn: func(_ context.Context, name, ownerID string) (*organization.Organization, error) {
				return &organization.Organization{ID: "id", Name: "name", OwnerID: "ownerid"}, nil
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
