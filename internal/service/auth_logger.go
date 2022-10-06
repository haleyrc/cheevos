package service

import (
	"context"

	"github.com/haleyrc/cheevos/lib/logger"
)

type Logger struct {
	Logger  logger.Logger
	Service interface {
		SignUp(ctx context.Context, username, password string) (*User, error)
	}
}

func (l *Logger) SignUp(ctx context.Context, username, password string) (*User, error) {
	l.Logger.Debug(ctx, "signing up user", logger.Fields{
		"Username": username,
	})

	user, err := l.Service.SignUp(ctx, username, password)
	if err != nil {
		l.Logger.Error(ctx, "sign up failed", err)
		return nil, err
	}

	l.Logger.Log(ctx, "user signed up", logger.Fields{
		"User": user,
	})

	return user, nil
}
