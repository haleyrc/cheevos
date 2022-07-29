package cheevos

import (
	"context"

	"github.com/haleyrc/cheevos/log"
)

type UserLogger struct {
	Svc interface {
		SignUp(ctx context.Context, req SignUpRequest) (*SignUpResponse, error)
	}
	Logger log.Logger
}

func (ol *UserLogger) SignUp(ctx context.Context, req SignUpRequest) (*SignUpResponse, error) {
	ol.Logger.Debug(ctx, "signing up user", log.Fields{"Username": req.Username})

	resp, err := ol.Svc.SignUp(ctx, req)
	if err != nil {
		ol.Logger.Error(ctx, "sign up failed", err)
		return nil, err
	}
	ol.Logger.Log(ctx, "user signed up", log.Fields{"User": resp.User})

	return resp, nil
}
