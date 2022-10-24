package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/haleyrc/pkg/errors"
	"github.com/haleyrc/pkg/hash"
	"github.com/haleyrc/pkg/logger"
	"github.com/haleyrc/pkg/pg"
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/domain"
)

var _ domain.AuthService = &authService{}

type AuthRepository interface {
	GetUser(ctx context.Context, tx pg.Tx, u *domain.User, id string) error
	InsertUser(ctx context.Context, tx pg.Tx, u *domain.User, hashedPassword string) error
}

func NewAuthService(db Database, logger logger.Logger, repo AuthRepository) domain.AuthService {
	return &authLogger{
		Logger: logger,
		Service: &authService{
			DB:   db,
			Repo: repo,
		},
	}
}

type authService struct {
	DB   Database
	Repo AuthRepository
}

func (svc *authService) GetUser(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User

	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx pg.Tx) error {
		return svc.Repo.GetUser(ctx, tx, &user, id)
	})
	if err != nil {
		return nil, fmt.Errorf("get user failed: %w", err)
	}

	return &user, nil
}

// SignUp creates a new user and persists it to the database. It returns a
// response containing the new organization if successful.
func (svc *authService) SignUp(ctx context.Context, username, password string) (*domain.User, error) {
	var user domain.User

	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx pg.Tx) error {
		user = domain.User{
			ID:       uuid.New(),
			Username: username,
		}
		if err := user.Validate(); err != nil {
			return errors.WrapError(err)
		}

		password = normalizePassword(password)
		if err := validatePassword(&user, password); err != nil {
			return errors.WrapError(err)
		}

		return svc.Repo.InsertUser(ctx, tx, &user, hash.Generate(password))
	})
	if err != nil {
		return nil, errors.WrapError(err)
	}

	return &user, nil
}

func normalizePassword(password string) string {
	return strings.TrimSpace(password)
}

// The User parameter here is required so we can construct our validation error
// correctly, but this feels like a pretty gnarly way of doing things.
func validatePassword(u *domain.User, password string) error {
	if len(password) < 8 {
		return domain.NewValidationError(u).
			Add("Password", "Password must be eighth (8) or more characters.").
			Error()
	}
	return nil
}
