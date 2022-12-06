package service

import (
	"context"
	"fmt"

	"github.com/haleyrc/pkg/errors"
	"github.com/haleyrc/pkg/logger"
	"github.com/haleyrc/pkg/pg"
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/domain"
	"github.com/haleyrc/cheevos/internal/password"
)

var _ domain.AuthService = &authService{}

type AuthRepository interface {
	GetUser(ctx context.Context, tx pg.Tx, u *domain.User, id string) error
	InsertUser(ctx context.Context, tx pg.Tx, u *domain.User, hashedPassword string) error
}

type PasswordValidator interface {
	Validate(p password.Password) error
}

func NewAuthService(db Database, logger logger.Logger, repo AuthRepository, pwv PasswordValidator) domain.AuthService {
	return &authLogger{
		Logger: logger,
		Service: &authService{
			DB:                db,
			Repo:              repo,
			PasswordValidator: pwv,
		},
	}
}

type authService struct {
	DB                Database
	Repo              AuthRepository
	PasswordValidator PasswordValidator
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
func (svc *authService) SignUp(ctx context.Context, username string, password password.Password) (*domain.User, error) {
	var user domain.User

	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx pg.Tx) error {
		user = domain.User{
			ID:       uuid.New(),
			Username: username,
		}
		if err := user.Validate(); err != nil {
			return errors.WrapError(err)
		}

		if err := svc.PasswordValidator.Validate(password); err != nil {
			return errors.WrapError(err)
		}

		return svc.Repo.InsertUser(ctx, tx, &user, password.Hash())
	})
	if err != nil {
		return nil, errors.WrapError(err)
	}

	return &user, nil
}
