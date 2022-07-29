package cheevos

import (
	"context"
)

type CheevoLogger struct {
	Svc interface {
		AwardCheevoToUser(context.Context, AwardCheevoToUserRequest) (*AwardCheevoToUserResponse, error)
		CreateCheevo(context.Context, CreateCheevoRequest) (*CreateCheevoResponse, error)
	}
	Logger Logger
}

func (cl *CheevoLogger) AwardCheevoToUser(ctx context.Context, req AwardCheevoToUserRequest) (*AwardCheevoToUserResponse, error) {
	cl.Logger.Debug(ctx, "awarding cheevo to user", Fields{
		"Cheevo":  req.Cheevo,
		"Awardee": req.Awardee,
		"Awarder": req.Awarder,
	})

	resp, err := cl.Svc.AwardCheevoToUser(ctx, req)
	if err != nil {
		cl.Logger.Error(ctx, "award cheevo to user failed", err)
		return nil, err
	}
	cl.Logger.Log(ctx, "cheevo awarded", Fields{
		"Cheevo": resp.Cheevo,
		"User":   resp.User,
	})

	return resp, nil
}

func (cl *CheevoLogger) CreateCheevo(ctx context.Context, req CreateCheevoRequest) (*CreateCheevoResponse, error) {
	cl.Logger.Debug(ctx, "creating cheevo", Fields{
		"Name":         req.Name,
		"Description":  req.Description,
		"Organization": req.Organization,
	})

	resp, err := cl.Svc.CreateCheevo(ctx, req)
	if err != nil {
		cl.Logger.Error(ctx, "create cheevo failed", err)
		return nil, err
	}
	cl.Logger.Log(ctx, "cheevo created", Fields{"Cheevo": resp.Cheevo})

	return resp, nil
}
