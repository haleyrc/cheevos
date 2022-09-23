package web

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos/cheevos"
	"github.com/haleyrc/cheevos/roster"
)

type AuthorizationService struct {
	Cheevos cheevos.Service
	Roster  roster.Service
}

func (svc *AuthorizationService) CanAwardCheevo(ctx context.Context, fromUserID, toUserID, cheevoID string) error {
	cheevo, err := svc.Cheevos.GetCheevo(ctx, cheevoID)
	if err != nil {
		return fmt.Errorf("authorization failed: %w", err)
	}

	if err := svc.Roster.IsMember(ctx, cheevo.OrganizationID, toUserID); err != nil {
		return fmt.Errorf("authorization failed: %w", err)
	}

	if err := svc.Roster.IsMember(ctx, cheevo.OrganizationID, fromUserID); err != nil {
		return fmt.Errorf("authorization failed: %w", err)
	}

	return nil
}

func (svc *AuthorizationService) CanCreateCheevo(ctx context.Context, userID, orgID string) error {
	if err := svc.Roster.IsMember(ctx, orgID, userID); err != nil {
		return fmt.Errorf("authorization failed: %w", err)
	}
	return nil
}

func (svc *AuthorizationService) CanGetCheevo(ctx context.Context, userID, cheevoID string) error {
	cheevo, err := svc.Cheevos.GetCheevo(ctx, cheevoID)
	if err != nil {
		return fmt.Errorf("authorization failed: %w", err)
	}

	if err := svc.Roster.IsMember(ctx, cheevo.OrganizationID, userID); err != nil {
		return fmt.Errorf("authorization failed: %w", err)
	}

	return nil
}
