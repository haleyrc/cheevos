package service

import (
	"context"

	"github.com/haleyrc/pkg/logger"

	"github.com/haleyrc/cheevos/domain"
	"github.com/haleyrc/cheevos/internal/password"
)

var _ domain.AuthService = &authLogger{}

type authLogger struct {
	Logger  logger.Logger
	Service domain.AuthService
}

func (l *authLogger) GetUser(ctx context.Context, id string) (*domain.User, error) {
	l.Logger.Debug(ctx, "getting user", logger.Fields{"ID": id})

	user, err := l.Service.GetUser(ctx, id)
	if err != nil {
		l.Logger.Error(ctx, "get user failed", err)
		return nil, err
	}

	l.Logger.Log(ctx, "got user", logger.Fields{"User": user})

	return user, nil
}

func (l *authLogger) SignUp(ctx context.Context, username string, password password.Password) (*domain.User, error) {
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
