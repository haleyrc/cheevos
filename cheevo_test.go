package cheevos_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/mock"
)

func TestCheevoLoggerLogsAnErrorFromCreateCheevo(t *testing.T) {
	logger := NewTestLogger()

	cl := &cheevos.CheevoLogger{
		Svc: &mock.CheevoService{
			CreateCheevoFn: func(_ context.Context, req cheevos.CreateCheevoRequest) (*cheevos.CreateCheevoResponse, error) {
				return nil, fmt.Errorf("oops")
			},
		},
		Logger: logger,
	}
	cl.CreateCheevo(context.Background(), cheevos.CreateCheevoRequest{
		Name:         "Test",
		Description:  "This is a test.",
		Organization: "783bf2de-dce2-4f32-9f18-f77b904f87c",
	})

	logger.ShouldLog(t,
		`{"Fields":{"Description":"This is a test.","Name":"Test","Organization":"783bf2de-dce2-4f32-9f18-f77b904f87c"},"Message":"creating cheevo"}`,
		`{"Fields":{"Error":"oops"},"Message":"create cheevo failed"}`,
	)
}

func TestCheevoLoggerLogsTheResponseFromCreateCheevo(t *testing.T) {
	logger := NewTestLogger()

	cl := &cheevos.CheevoLogger{
		Svc: &mock.CheevoService{
			CreateCheevoFn: func(_ context.Context, req cheevos.CreateCheevoRequest) (*cheevos.CreateCheevoResponse, error) {
				return &cheevos.CreateCheevoResponse{
					Cheevo: &cheevos.Cheevo{
						ID:           "8059dcd7-bcc1-46fa-bfc0-3926c0b2c6ea",
						Name:         "Test",
						Description:  "This is a test.",
						Organization: "238cb95f-8bcd-4cda-8cfc-9d03fecba894",
					},
				}, nil
			},
		},
		Logger: logger,
	}
	cl.CreateCheevo(context.Background(), cheevos.CreateCheevoRequest{
		Name:         "Test",
		Description:  "This is a test.",
		Organization: "238cb95f-8bcd-4cda-8cfc-9d03fecba894",
	})

	logger.ShouldLog(t,
		`{"Fields":{"Description":"This is a test.","Name":"Test","Organization":"238cb95f-8bcd-4cda-8cfc-9d03fecba894"},"Message":"creating cheevo"}`,
		`{"Fields":{"Cheevo":{"ID":"8059dcd7-bcc1-46fa-bfc0-3926c0b2c6ea","Name":"Test","Description":"This is a test.","Organization":"238cb95f-8bcd-4cda-8cfc-9d03fecba894"}},"Message":"cheevo created"}`,
	)
}

func TestCheevoLoggerLogsAnErrorFromAwardCheevoToUser(t *testing.T) {
	logger := NewTestLogger()

	cl := &cheevos.CheevoLogger{
		Svc: &mock.CheevoService{
			AwardCheevoToUserFn: func(_ context.Context, req cheevos.AwardCheevoToUserRequest) (*cheevos.AwardCheevoToUserResponse, error) {
				return nil, fmt.Errorf("oops")
			},
		},
		Logger: logger,
	}
	cl.AwardCheevoToUser(context.Background(), cheevos.AwardCheevoToUserRequest{
		Cheevo: "783bf2de-dce2-4f32-9f18-f77b904f87c",
		User:   "4d523938-2baa-4d94-8daf-ea1785ff154",
	})

	logger.ShouldLog(t,
		`{"Fields":{"Cheevo":"783bf2de-dce2-4f32-9f18-f77b904f87c","User":"4d523938-2baa-4d94-8daf-ea1785ff154"},"Message":"awarding cheevo to user"}`,
		`{"Fields":{"Error":"oops"},"Message":"award cheevo to user failed"}`,
	)
}

func TestCheevoLoggerLogsTheResponseFromAwardCheevoToUser(t *testing.T) {
	logger := NewTestLogger()

	cl := &cheevos.CheevoLogger{
		Svc: &mock.CheevoService{
			AwardCheevoToUserFn: func(_ context.Context, req cheevos.AwardCheevoToUserRequest) (*cheevos.AwardCheevoToUserResponse, error) {
				return &cheevos.AwardCheevoToUserResponse{
					Cheevo: &cheevos.Cheevo{
						ID:           "783bf2de-dce2-4f32-9f18-f77b904f87cf",
						Name:         "Test",
						Description:  "This is a test.",
						Organization: "4d523938-2baa-4d94-8daf-ea1785ff154d",
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
	cl.AwardCheevoToUser(context.Background(), cheevos.AwardCheevoToUserRequest{
		Cheevo: "783bf2de-dce2-4f32-9f18-f77b904f87cf",
		User:   "4d523938-2baa-4d94-8daf-ea1785ff154d",
	})

	logger.ShouldLog(t,
		`{"Fields":{"Cheevo":"783bf2de-dce2-4f32-9f18-f77b904f87cf","User":"4d523938-2baa-4d94-8daf-ea1785ff154d"},"Message":"awarding cheevo to user"}`,
		`{"Fields":{"Cheevo":{"ID":"783bf2de-dce2-4f32-9f18-f77b904f87cf","Name":"Test","Description":"This is a test.","Organization":"4d523938-2baa-4d94-8daf-ea1785ff154d"},"User":{"ID":"2d7c6d16-c703-4058-a4dd-fb8d34992806","Username":"test"}},"Message":"cheevo awarded"}`,
	)
}

func TestAwardingACheevoToAUserSucceeds(t *testing.T) {
	ctx := context.Background()
	db := mock.NewDatabase()
	db.AwardCheevoToUserFn = func(_ context.Context, _, _ string) error {
		return nil
	}
	db.GetCheevoFn = func(_ context.Context, cheevoID string) (*cheevos.Cheevo, error) {
		return &cheevos.Cheevo{ID: cheevoID, Name: "Test", Description: "Test cheevo."}, nil
	}
	db.GetUserFn = func(_ context.Context, userID string) (*cheevos.User, error) {
		return &cheevos.User{ID: userID, Username: "test"}, nil
	}
	svc := cheevos.CheevoService{DB: db}
	cheevoID := uuid.New()
	userID := uuid.New()

	resp, err := svc.AwardCheevoToUser(ctx, cheevos.AwardCheevoToUserRequest{
		Cheevo: cheevoID,
		User:   userID,
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Cheevo.ID != cheevoID {
		t.Errorf("Cheevo should be %q, but got %q.", cheevoID, resp.Cheevo.ID)
	}
	if resp.User.ID != userID {
		t.Errorf("User should be %q, but got %q.", userID, resp.User.ID)
	}
}

func TestCreatingAValidCheevoWithSucceeds(t *testing.T) {
	ctx := context.Background()
	db := mock.NewDatabase()
	svc := cheevos.CheevoService{DB: db}
	orgID := uuid.New()

	resp, err := svc.CreateCheevo(ctx, cheevos.CreateCheevoRequest{
		Name:         "Test",
		Description:  "This is a test cheevo.",
		Organization: orgID,
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Cheevo.ID == "" {
		t.Error("ID shouldn't be blank, but it was.")
	}
	if resp.Cheevo.Name != "Test" {
		t.Errorf("Name should be \"Test\", but got %q.", resp.Cheevo.Name)
	}
	if resp.Cheevo.Description != "This is a test cheevo." {
		t.Errorf("Description should be \"This is a test cheevo.\", but got %q.", resp.Cheevo.Description)
	}
	if resp.Cheevo.Organization != orgID {
		t.Errorf("Organization should be %q, but got %q.", orgID, resp.Cheevo.Organization)
	}
}

func TestCreatingACheevoWithAnInvalidOrganizationFails(t *testing.T) {
	ctx := context.Background()
	svc := cheevos.CheevoService{}

	// We don't have to test blank org here for the same reason we don't have to
	// normalize it: the org ID is not provided by the user so the ID will either
	// exist or it won't.
	testcases := map[string]string{
		"empty org": "",
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			_, err := svc.CreateCheevo(ctx, cheevos.CreateCheevoRequest{
				Name:         "test",
				Description:  "testtest",
				Organization: tc,
			})
			if err == nil {
				t.Error("expected an error, but got none")
			}
		})
	}
}

func TestCreatingACheevoWithAnInvalidNameFails(t *testing.T) {
	ctx := context.Background()
	svc := cheevos.CheevoService{}

	testcases := map[string]string{
		"empty name": "",
		"blank name": " \t\n",
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			_, err := svc.CreateCheevo(ctx, cheevos.CreateCheevoRequest{
				Name:         tc,
				Description:  "testtest",
				Organization: uuid.New(),
			})
			if err == nil {
				t.Error("expected an error, but got none")
			}
		})
	}
}

func TestCreatingACheevoWithAnInvalidDescriptionFails(t *testing.T) {
	ctx := context.Background()
	svc := cheevos.CheevoService{}

	testcases := map[string]string{
		"empty description": "",
		"blank description": " \t\n",
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			_, err := svc.CreateCheevo(ctx, cheevos.CreateCheevoRequest{
				Name:         "Test",
				Description:  tc,
				Organization: uuid.New(),
			})
			if err == nil {
				t.Error("expected an error, but got none")
			}
		})
	}
}
