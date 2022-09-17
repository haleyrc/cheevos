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

func TestOrganizationLoggerLogsAnErrorFromAddUserToOrganization(t *testing.T) {
	logger := testutil.NewTestLogger()

	ol := &organization.OrganizationLogger{
		Svc: &mock.OrganizationService{
			AddMemberToOrganizationFn: func(_ context.Context, _, _ string) error {
				return fmt.Errorf("oops")
			},
		},
		Logger: logger,
	}
	ol.AddMemberToOrganization(context.Background(), "4d523938-2baa-4d94-8daf-ea1785ff154", "783bf2de-dce2-4f32-9f18-f77b904f87c")

	logger.ShouldLog(t,
		`{"Fields":{"Organization":"783bf2de-dce2-4f32-9f18-f77b904f87c","User":"4d523938-2baa-4d94-8daf-ea1785ff154"},"Message":"adding member to organization"}`,
		`{"Fields":{"Error":"oops"},"Message":"add member to organization failed"}`,
	)
}

func TestOrganizationLoggerLogsTheResponseFromAddUserToOrganization(t *testing.T) {
	logger := testutil.NewTestLogger()

	ol := &organization.OrganizationLogger{
		Svc: &mock.OrganizationService{
			AddMemberToOrganizationFn: func(_ context.Context, userID, ownerID string) error { return nil },
		},
		Logger: logger,
	}
	ol.AddMemberToOrganization(context.Background(), "4d523938-2baa-4d94-8daf-ea1785ff154d", "783bf2de-dce2-4f32-9f18-f77b904f87cf")

	logger.ShouldLog(t,
		`{"Fields":{"Organization":"783bf2de-dce2-4f32-9f18-f77b904f87cf","User":"4d523938-2baa-4d94-8daf-ea1785ff154d"},"Message":"adding member to organization"}`,
		`{"Fields":{"Organization":"783bf2de-dce2-4f32-9f18-f77b904f87cf","User":"4d523938-2baa-4d94-8daf-ea1785ff154d"},"Message":"added member to organization"}`,
	)
}

func TestOrganizationLoggerLogsAnErrorFromCreateOrganization(t *testing.T) {
	logger := testutil.NewTestLogger()

	ol := &organization.OrganizationLogger{
		Svc: &mock.OrganizationService{
			CreateOrganizationFn: func(_ context.Context, name, ownerID string) (*organization.Organization, error) {
				return nil, fmt.Errorf("oops")
			},
		},
		Logger: logger,
	}
	ol.CreateOrganization(context.Background(), "Test", "783bf2de-dce2-4f32-9f18-f77b904f87c")

	logger.ShouldLog(t,
		`{"Fields":{"Name":"Test","Owner":"783bf2de-dce2-4f32-9f18-f77b904f87c"},"Message":"creating organization"}`,
		`{"Fields":{"Error":"oops"},"Message":"create organization failed"}`,
	)
}

func TestOrganizationLoggerLogsTheResponseFromCreateOrganization(t *testing.T) {
	time.Freeze()

	logger := testutil.NewTestLogger()

	ol := &organization.OrganizationLogger{
		Svc: &mock.OrganizationService{
			CreateOrganizationFn: func(_ context.Context, name, ownerID string) (*organization.Organization, error) {
				return &organization.Organization{
					ID:      "8059dcd7-bcc1-46fa-bfc0-3926c0b2c6ea",
					Name:    "Test",
					OwnerID: "238cb95f-8bcd-4cda-8cfc-9d03fecba894",
				}, nil
			},
		},
		Logger: logger,
	}
	ol.CreateOrganization(context.Background(), "Test", "238cb95f-8bcd-4cda-8cfc-9d03fecba894")

	logger.ShouldLog(t,
		`{"Fields":{"Name":"Test","Owner":"238cb95f-8bcd-4cda-8cfc-9d03fecba894"},"Message":"creating organization"}`,
		`{"Fields":{"Organization":{"ID":"8059dcd7-bcc1-46fa-bfc0-3926c0b2c6ea","Name":"Test","OwnerID":"238cb95f-8bcd-4cda-8cfc-9d03fecba894"}},"Message":"organization created"}`,
	)
}
