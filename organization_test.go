package cheevos_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/mock"
)

func TestOrganizationLoggerLogsAnErrorFromAddUserToOrganization(t *testing.T) {
	logger := NewTestLogger()

	cl := &cheevos.OrganizationLogger{
		Svc: &mock.OrganizationService{
			AddUserToOrganizationFn: func(_ context.Context, req cheevos.AddUserToOrganizationRequest) (*cheevos.AddUserToOrganizationResponse, error) {
				return nil, fmt.Errorf("oops")
			},
		},
		Logger: logger,
	}
	cl.AddUserToOrganization(context.Background(), cheevos.AddUserToOrganizationRequest{
		Organization: "783bf2de-dce2-4f32-9f18-f77b904f87c",
		User:         "4d523938-2baa-4d94-8daf-ea1785ff154",
	})

	logger.ShouldLog(t,
		`{"Fields":{"Organization":"783bf2de-dce2-4f32-9f18-f77b904f87c","User":"4d523938-2baa-4d94-8daf-ea1785ff154"},"Message":"adding user to organization"}`,
		`{"Fields":{"Error":"oops"},"Message":"add user to organization failed"}`,
	)
}

func TestOrganizationLoggerLogsTheResponseFromAddUserToOrganization(t *testing.T) {
	logger := NewTestLogger()

	cl := &cheevos.OrganizationLogger{
		Svc: &mock.OrganizationService{
			AddUserToOrganizationFn: func(_ context.Context, req cheevos.AddUserToOrganizationRequest) (*cheevos.AddUserToOrganizationResponse, error) {
				return &cheevos.AddUserToOrganizationResponse{
					Organization: &cheevos.Organization{
						ID:    "783bf2de-dce2-4f32-9f18-f77b904f87cf",
						Name:  "Test",
						Owner: "4d523938-2baa-4d94-8daf-ea1785ff154d",
					},
					User: &cheevos.User{
						ID:       "2d7c6d16-c703-4058-a4dd-fb8d34992806",
						Username: "test",
					},
				}, nil
			},
		},
		Logger: logger,
	}
	cl.AddUserToOrganization(context.Background(), cheevos.AddUserToOrganizationRequest{
		Organization: "783bf2de-dce2-4f32-9f18-f77b904f87cf",
		User:         "4d523938-2baa-4d94-8daf-ea1785ff154d",
	})

	logger.ShouldLog(t,
		`{"Fields":{"Organization":"783bf2de-dce2-4f32-9f18-f77b904f87cf","User":"4d523938-2baa-4d94-8daf-ea1785ff154d"},"Message":"adding user to organization"}`,
		`{"Fields":{"Organization":{"ID":"783bf2de-dce2-4f32-9f18-f77b904f87cf","Name":"Test","Owner":"4d523938-2baa-4d94-8daf-ea1785ff154d"},"User":{"ID":"2d7c6d16-c703-4058-a4dd-fb8d34992806","Username":"test"}},"Message":"user added to organization"}`,
	)
}

func TestOrganizationLoggerLogsAnErrorFromCreateOrganization(t *testing.T) {
	logger := NewTestLogger()

	cl := &cheevos.OrganizationLogger{
		Svc: &mock.OrganizationService{
			CreateOrganizationFn: func(_ context.Context, req cheevos.CreateOrganizationRequest) (*cheevos.CreateOrganizationResponse, error) {
				return nil, fmt.Errorf("oops")
			},
		},
		Logger: logger,
	}
	cl.CreateOrganization(context.Background(), cheevos.CreateOrganizationRequest{
		Name:  "Test",
		Owner: "783bf2de-dce2-4f32-9f18-f77b904f87c",
	})

	logger.ShouldLog(t,
		`{"Fields":{"Name":"Test","Owner":"783bf2de-dce2-4f32-9f18-f77b904f87c"},"Message":"creating organization"}`,
		`{"Fields":{"Error":"oops"},"Message":"create organization failed"}`,
	)
}

func TestOrganizationLoggerLogsTheResponseFromCreateOrganization(t *testing.T) {
	logger := NewTestLogger()

	cl := &cheevos.OrganizationLogger{
		Svc: &mock.OrganizationService{
			CreateOrganizationFn: func(_ context.Context, req cheevos.CreateOrganizationRequest) (*cheevos.CreateOrganizationResponse, error) {
				return &cheevos.CreateOrganizationResponse{
					Organization: &cheevos.Organization{
						ID:    "8059dcd7-bcc1-46fa-bfc0-3926c0b2c6ea",
						Name:  "Test",
						Owner: "238cb95f-8bcd-4cda-8cfc-9d03fecba894",
					},
				}, nil
			},
		},
		Logger: logger,
	}
	cl.CreateOrganization(context.Background(), cheevos.CreateOrganizationRequest{
		Name:  "Test",
		Owner: "238cb95f-8bcd-4cda-8cfc-9d03fecba894",
	})

	logger.ShouldLog(t,
		`{"Fields":{"Name":"Test","Owner":"238cb95f-8bcd-4cda-8cfc-9d03fecba894"},"Message":"creating organization"}`,
		`{"Fields":{"Organization":{"ID":"8059dcd7-bcc1-46fa-bfc0-3926c0b2c6ea","Name":"Test","Owner":"238cb95f-8bcd-4cda-8cfc-9d03fecba894"}},"Message":"organization created"}`,
	)
}

func TestAddingAUserToAnOrganizationSucceeds(t *testing.T) {
	ctx := context.Background()
	db := mock.NewDatabase()
	db.AddUserToOrganizationFn = func(ctx context.Context, orgID, userID string) error {
		return nil
	}
	db.GetOrganizationFn = func(ctx context.Context, orgID string) (*cheevos.Organization, error) {
		return &cheevos.Organization{ID: orgID, Name: "Test"}, nil
	}
	db.GetUserFn = func(ctx context.Context, userID string) (*cheevos.User, error) {
		return &cheevos.User{ID: userID, Username: "test"}, nil
	}
	svc := cheevos.OrganizationService{DB: db}
	orgID := uuid.New()
	userID := uuid.New()

	resp, err := svc.AddUserToOrganization(ctx, cheevos.AddUserToOrganizationRequest{
		Organization: orgID,
		User:         userID,
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Organization.ID != orgID {
		t.Errorf("Organization should be %q, but got %q.", orgID, resp.Organization.ID)
	}
	if resp.User.ID != userID {
		t.Errorf("User should be %q, but got %q.", userID, resp.User.ID)
	}
}

func TestCreatingAValidOrganizationWithSucceeds(t *testing.T) {
	ctx := context.Background()
	db := mock.NewDatabase()
	svc := cheevos.OrganizationService{DB: db}
	ownerID := uuid.New()

	resp, err := svc.CreateOrganization(ctx, cheevos.CreateOrganizationRequest{
		Name:  "Test",
		Owner: ownerID,
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Organization.ID == "" {
		t.Error("ID shouldn't be blank, but it was.")
	}
	if resp.Organization.Name != "Test" {
		t.Errorf("Name should be \"Test\", but got %q.", resp.Organization.Name)
	}
	if resp.Organization.Owner != ownerID {
		t.Errorf("Owner should be %q, but got %q.", ownerID, resp.Organization.Owner)
	}
}

func TestCreatingAOrganizationWithAnInvalidNameFails(t *testing.T) {
	ctx := context.Background()
	svc := cheevos.OrganizationService{}

	testcases := map[string]string{
		"empty name": "",
		"blank name": " \t\n",
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			_, err := svc.CreateOrganization(ctx, cheevos.CreateOrganizationRequest{
				Name:  tc,
				Owner: uuid.New(),
			})
			if err == nil {
				t.Error("expected an error, but got none")
			}
		})
	}
}

func TestCreatingAOrganizationWithAnInvalidOwnerFails(t *testing.T) {
	ctx := context.Background()
	svc := cheevos.OrganizationService{}

	// We don't have to test for a blank owner here for the same reason we don't
	// have to normalize it: it's not coming from the user so it either exists or
	// it doesn't.
	testcases := map[string]string{
		"empty name": "",
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			_, err := svc.CreateOrganization(ctx, cheevos.CreateOrganizationRequest{
				Name:  "test",
				Owner: tc,
			})
			if err == nil {
				t.Error("expected an error, but got none")
			}
		})
	}
}
