package mock

import "context"

type AwardService struct {
	AwardCheevoToUserFn func(ctx context.Context, recipientID, cheevoID string) error
}

func (as *AwardService) AwardCheevoToUser(ctx context.Context, recipientID, cheevoID string) error {
	return as.AwardCheevoToUserFn(ctx, recipientID, cheevoID)
}
