package mock

import "context"

type MembershipService struct {
	AddMemberToOrganizationFn func(ctx context.Context, userID, orgID string) error
}

func (ms *MembershipService) AddMemberToOrganization(ctx context.Context, userID, orgID string) error {
	return ms.AddMemberToOrganizationFn(ctx, userID, orgID)
}
