package mock

import (
	"context"

	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/membership"
)

type CreateMembershipArgs struct {
	Membership *membership.Membership
}

type MembershipRepository struct {
	CreateMembershipFn     func(ctx context.Context, tx db.Transaction, m *membership.Membership) error
	CreateMembershipCalled struct {
		Count int
		With  CreateMembershipArgs
	}
}

func (mr *MembershipRepository) CreateMembership(ctx context.Context, tx db.Transaction, m *membership.Membership) error {
	if mr.CreateMembershipFn == nil {
		return mockMethodNotDefined("CreateMembership")
	}
	mr.CreateMembershipCalled.Count++
	mr.CreateMembershipCalled.With = CreateMembershipArgs{Membership: m}
	return mr.CreateMembershipFn(ctx, tx, m)
}
