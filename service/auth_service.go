package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/haleyrc/pkg/hash"
	"github.com/haleyrc/pkg/logger"
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/core"
	"github.com/haleyrc/cheevos/internal/lib/db"
)

var _ cheevos.AuthService = &authService{}

type AuthRepository interface {
	GetUser(ctx context.Context, tx db.Tx, u *cheevos.User, id string) error
	InsertUser(ctx context.Context, tx db.Tx, u *cheevos.User, hashedPassword string) error
}

func NewAuthService(db db.Database, logger logger.Logger, repo AuthRepository) cheevos.AuthService {
	return &authLogger{
		Logger: logger,
		Service: &authService{
			DB:   db,
			Repo: repo,
		},
	}
}

type authService struct {
	DB   db.Database
	Repo AuthRepository
}

func (svc *authService) GetUser(ctx context.Context, id string) (*cheevos.User, error) {
	var user cheevos.User

	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx db.Tx) error {
		return svc.Repo.GetUser(ctx, tx, &user, id)
	})
	if err != nil {
		return nil, fmt.Errorf("get user failed: %w", err)
	}

	return &user, nil
}

// SignUp creates a new user and persists it to the database. It returns a
// response containing the new organization if successful.
func (svc *authService) SignUp(ctx context.Context, username, password string) (*cheevos.User, error) {
	var user cheevos.User

	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx db.Tx) error {
		user = cheevos.User{
			ID:       uuid.New(),
			Username: username,
		}
		if err := user.Validate(); err != nil {
			return core.WrapError(err)
		}

		password = normalizePassword(password)
		if err := validatePassword(&user, password); err != nil {
			return core.WrapError(err)
		}

		return svc.Repo.InsertUser(ctx, tx, &user, hash.Generate(password))
	})
	if err != nil {
		return nil, core.WrapError(err)
	}

	return &user, nil
}

func normalizePassword(password string) string {
	return strings.TrimSpace(password)
}

// The User parameter here is required so we can construct our validation error
// correctly, but this feels like a pretty gnarly way of doing things.
func validatePassword(u *cheevos.User, password string) error {
	if len(password) < 8 {
		return core.NewValidationError(u).
			Add("Password", "Password must be eighth (8) or more characters.").
			Error()
	}
	return nil
}
