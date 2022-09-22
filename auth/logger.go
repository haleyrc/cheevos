package auth

import (
	"context"

	"github.com/haleyrc/cheevos/lib/logger"
)

type Logger struct {
	Svc interface {
		SignUp(ctx context.Context, username, password string) (*User, error)
	}
	Logger logger.Logger
}

func (ul *Logger) SignUp(ctx context.Context, username, password string) (*User, error) {
	ul.Logger.Debug(ctx, "signing up user", logger.Fields{
		"Username": username,
	})

	user, err := ul.Svc.SignUp(ctx, username, password)
	if err != nil {
		ul.Logger.Error(ctx, "sign up failed", err)
		return nil, err
	}

	ul.Logger.Log(ctx, "user signed up", logger.Fields{
		"User": user,
	})

	return user, nil
}
