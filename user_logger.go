package cheevos

import "context"

type UserLogger struct {
	Svc interface {
		SignUp(ctx context.Context, req SignUpRequest) (*SignUpResponse, error)
	}
	Logger Logger
}

func (ol *UserLogger) SignUp(ctx context.Context, req SignUpRequest) (*SignUpResponse, error) {
	ol.Logger.Debug(ctx, "signing up user", Fields{"Username": req.Username})

	resp, err := ol.Svc.SignUp(ctx, req)
	if err != nil {
		ol.Logger.Error(ctx, "sign up failed", err)
		return nil, err
	}
	ol.Logger.Log(ctx, "user signed up", Fields{"User": resp.User})

	return resp, nil
}
