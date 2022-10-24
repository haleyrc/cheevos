package mock

import (
	"context"
	"fmt"

	"github.com/haleyrc/pkg/pg"

	"github.com/haleyrc/cheevos/domain"
)

type DeleteInvitationByCodeArgs struct {
	Code string
}

type GetCheevoArgs struct {
	ID string
}

type GetInvitationArgs struct {
	ID string
}

type GetInvitationByCodeArgs struct {
	Code string
}

type GetMembershipArgs struct {
	OrganizationID string
	UserID         string
}

type GetUserArgs struct {
	ID string
}

type InsertAwardArgs struct {
	Award *domain.Award
}

type InsertCheevoArgs struct {
	Cheevo *domain.Cheevo
}

type InsertInvitationArgs struct {
	Invitation *domain.Invitation
	HashedCode string
}

type InsertMembershipArgs struct {
	Membership *domain.Membership
}

type InsertOrganizationArgs struct {
	Organization *domain.Organization
}

type InsertUserArgs struct {
	User         *domain.User
	PasswordHash string
}

type UpdateInvitationArgs struct {
	Invitation *domain.Invitation
	HashedCode string
}

type Repository struct {
	DeleteInvitationByCodeFn     func(ctx context.Context, tx pg.Tx, code string) error
	DeleteInvitationByCodeCalled struct {
		Count int
		With  DeleteInvitationByCodeArgs
	}

	GetCheevoFn     func(ctx context.Context, tx pg.Tx, cheevo *domain.Cheevo, id string) error
	GetCheevoCalled struct {
		Count int
		With  GetCheevoArgs
	}

	GetInvitationFn     func(ctx context.Context, tx pg.Tx, i *domain.Invitation, id string) error
	GetInvitationCalled struct {
		Count int
		With  GetInvitationArgs
	}

	GetInvitationByCodeFn     func(ctx context.Context, tx pg.Tx, i *domain.Invitation, code string) error
	GetInvitationByCodeCalled struct {
		Count int
		With  GetInvitationByCodeArgs
	}

	GetMembershipFn     func(ctx context.Context, tx pg.Tx, m *domain.Membership, orgID, userID string) error
	GetMembershipCalled struct {
		Count int
		With  GetMembershipArgs
	}

	GetUserFn     func(ctx context.Context, tx pg.Tx, u *domain.User, id string) error
	GetUserCalled struct {
		Count int
		With  GetUserArgs
	}

	InsertAwardFn     func(ctx context.Context, tx pg.Tx, a *domain.Award) error
	InsertAwardCalled struct {
		Count int
		With  InsertAwardArgs
	}

	InsertCheevoFn     func(ctx context.Context, tx pg.Tx, cheevo *domain.Cheevo) error
	InsertCheevoCalled struct {
		Count int
		With  InsertCheevoArgs
	}

	InsertInvitationFn     func(ctx context.Context, tx pg.Tx, i *domain.Invitation, hashedCode string) error
	InsertInvitationCalled struct {
		Count int
		With  InsertInvitationArgs
	}

	InsertMembershipFn     func(ctx context.Context, tx pg.Tx, m *domain.Membership) error
	InsertMembershipCalled struct {
		Count int
		With  InsertMembershipArgs
	}

	InsertOrganizationFn     func(ctx context.Context, tx pg.Tx, org *domain.Organization) error
	InsertOrganizationCalled struct {
		Count int
		With  InsertOrganizationArgs
	}

	InsertUserFn     func(ctx context.Context, tx pg.Tx, u *domain.User, hashedPassword string) error
	InsertUserCalled struct {
		Count int
		With  InsertUserArgs
	}

	UpdateInvitationFn     func(ctx context.Context, tx pg.Tx, i *domain.Invitation, hashedCode string) error
	UpdateInvitationCalled struct {
		Count int
		With  UpdateInvitationArgs
	}
}

func (repo *Repository) InsertAward(ctx context.Context, tx pg.Tx, a *domain.Award) error {
	if repo.InsertAwardFn == nil {
		return mockMethodNotDefined("InsertAward")
	}
	repo.InsertAwardCalled.Count++
	repo.InsertAwardCalled.With = InsertAwardArgs{Award: a}
	return repo.InsertAwardFn(ctx, tx, a)
}

func (repo *Repository) InsertCheevo(ctx context.Context, tx pg.Tx, cheevo *domain.Cheevo) error {
	if repo.InsertCheevoFn == nil {
		return mockMethodNotDefined("InsertCheevo")
	}
	repo.InsertCheevoCalled.Count++
	repo.InsertCheevoCalled.With = InsertCheevoArgs{Cheevo: cheevo}
	return repo.InsertCheevoFn(ctx, tx, cheevo)
}

func (repo *Repository) InsertInvitation(ctx context.Context, tx pg.Tx, i *domain.Invitation, hashedCode string) error {
	if repo.InsertInvitationFn == nil {
		return mockMethodNotDefined("InsertInvitation")
	}
	repo.InsertInvitationCalled.Count++
	repo.InsertInvitationCalled.With = InsertInvitationArgs{Invitation: i, HashedCode: hashedCode}
	return repo.InsertInvitationFn(ctx, tx, i, hashedCode)
}

func (repo *Repository) InsertMembership(ctx context.Context, tx pg.Tx, m *domain.Membership) error {
	if repo.InsertMembershipFn == nil {
		return mockMethodNotDefined("InsertMembership")
	}
	repo.InsertMembershipCalled.Count++
	repo.InsertMembershipCalled.With = InsertMembershipArgs{Membership: m}
	return repo.InsertMembershipFn(ctx, tx, m)
}

func (repo *Repository) InsertOrganization(ctx context.Context, tx pg.Tx, org *domain.Organization) error {
	if repo.InsertOrganizationFn == nil {
		return mockMethodNotDefined("InsertOrganization")
	}
	repo.InsertOrganizationCalled.Count++
	repo.InsertOrganizationCalled.With = InsertOrganizationArgs{Organization: org}
	return repo.InsertOrganizationFn(ctx, tx, org)
}

func (repo *Repository) InsertUser(ctx context.Context, tx pg.Tx, u *domain.User, hashedPassword string) error {
	if repo.InsertUserFn == nil {
		return mockMethodNotDefined("InsertUser")
	}
	repo.InsertUserCalled.Count++
	repo.InsertUserCalled.With = InsertUserArgs{User: u, PasswordHash: hashedPassword}
	return repo.InsertUserFn(ctx, tx, u, hashedPassword)
}

func (repo *Repository) DeleteInvitationByCode(ctx context.Context, tx pg.Tx, code string) error {
	if repo.DeleteInvitationByCodeFn == nil {
		return mockMethodNotDefined("DeleteInvitationByCode")
	}
	repo.DeleteInvitationByCodeCalled.Count++
	repo.DeleteInvitationByCodeCalled.With = DeleteInvitationByCodeArgs{Code: code}
	return repo.DeleteInvitationByCodeFn(ctx, tx, code)
}

func (repo *Repository) GetCheevo(ctx context.Context, tx pg.Tx, cheevo *domain.Cheevo, id string) error {
	if repo.GetCheevoFn == nil {
		return mockMethodNotDefined("GetCheevo")
	}
	repo.GetCheevoCalled.Count++
	repo.GetCheevoCalled.With = GetCheevoArgs{ID: id}
	return repo.GetCheevoFn(ctx, tx, cheevo, id)
}

func (repo *Repository) GetInvitation(ctx context.Context, tx pg.Tx, i *domain.Invitation, id string) error {
	if repo.GetInvitationFn == nil {
		return mockMethodNotDefined("GetInvitation")
	}
	repo.GetInvitationCalled.Count++
	repo.GetInvitationCalled.With = GetInvitationArgs{ID: id}
	return repo.GetInvitationFn(ctx, tx, i, id)
}

func (repo *Repository) GetInvitationByCode(ctx context.Context, tx pg.Tx, i *domain.Invitation, code string) error {
	if repo.GetInvitationByCodeFn == nil {
		return mockMethodNotDefined("GetInvitationByCode")
	}
	repo.GetInvitationByCodeCalled.Count++
	repo.GetInvitationByCodeCalled.With = GetInvitationByCodeArgs{Code: code}
	return repo.GetInvitationByCodeFn(ctx, tx, i, code)
}

func (repo *Repository) GetMembership(ctx context.Context, tx pg.Tx, m *domain.Membership, orgID, userID string) error {
	if repo.GetMembershipFn == nil {
		return mockMethodNotDefined("GetMembership")
	}
	repo.GetMembershipCalled.Count++
	repo.GetMembershipCalled.With = GetMembershipArgs{OrganizationID: orgID, UserID: userID}
	return repo.GetMembershipFn(ctx, tx, m, orgID, userID)
}

func (repo *Repository) GetUser(ctx context.Context, tx pg.Tx, u *domain.User, id string) error {
	if repo.GetUserFn == nil {
		return mockMethodNotDefined("GetUser")
	}
	repo.GetUserCalled.Count++
	repo.GetUserCalled.With = GetUserArgs{ID: id}
	return repo.GetUserFn(ctx, tx, u, id)
}

func (repo *Repository) UpdateInvitation(ctx context.Context, tx pg.Tx, i *domain.Invitation, hashedCode string) error {
	if repo.UpdateInvitationFn == nil {
		return mockMethodNotDefined("UpdateInvitation")
	}
	repo.UpdateInvitationCalled.Count++
	repo.UpdateInvitationCalled.With = UpdateInvitationArgs{Invitation: i, HashedCode: hashedCode}
	return repo.UpdateInvitationFn(ctx, tx, i, hashedCode)
}

func mockMethodNotDefined(funcName string) error {
	return fmt.Errorf("mock method %s is not defined", funcName)
}
