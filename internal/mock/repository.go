package mock

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos/award"
	"github.com/haleyrc/cheevos/cheevo"
	"github.com/haleyrc/cheevos/invitation"
	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/membership"
	"github.com/haleyrc/cheevos/organization"
	"github.com/haleyrc/cheevos/user"
)

type AddMemberToOrganizationArgs struct {
	OrganizationID string
	UserID         string
}

type CreateAwardArgs struct {
	Award *award.Award
}

type CreateCheevoArgs struct {
	Cheevo *cheevo.Cheevo
}

type CreateInvitationArgs struct {
	Invitation *invitation.Invitation
	HashedCode string
}

type CreateMembershipArgs struct {
	Membership *membership.Membership
}

type CreateOrganizationArgs struct {
	Organization *organization.Organization
}

type CreateUserArgs struct {
	User         *user.User
	PasswordHash string
}

type DeleteInvitationByCodeArgs struct {
	Code string
}

type GetInvitationByCodeArgs struct {
	Code string
}

type SaveInvitationArgs struct {
	Invitation *invitation.Invitation
	HashedCode string
}

type Repository struct {
	AddMemberToOrganizationFn     func(ctx context.Context, tx db.Transaction, userID, orgID string) error
	AddMemberToOrganizationCalled struct {
		Count int
		With  AddMemberToOrganizationArgs
	}

	CreateAwardFn     func(ctx context.Context, tx db.Transaction, a *award.Award) error
	CreateAwardCalled struct {
		Count int
		With  CreateAwardArgs
	}

	CreateCheevoFn     func(ctx context.Context, tx db.Transaction, cheevo *cheevo.Cheevo) error
	CreateCheevoCalled struct {
		Count int
		With  CreateCheevoArgs
	}

	CreateInvitationFn     func(ctx context.Context, tx db.Transaction, i *invitation.Invitation, hashedCode string) error
	CreateInvitationCalled struct {
		Count int
		With  CreateInvitationArgs
	}

	CreateMembershipFn     func(ctx context.Context, tx db.Transaction, m *membership.Membership) error
	CreateMembershipCalled struct {
		Count int
		With  CreateMembershipArgs
	}

	CreateOrganizationFn     func(ctx context.Context, tx db.Transaction, org *organization.Organization) error
	CreateOrganizationCalled struct {
		Count int
		With  CreateOrganizationArgs
	}

	CreateUserFn     func(ctx context.Context, tx db.Transaction, u *user.User, hashedPassword string) error
	CreateUserCalled struct {
		Count int
		With  CreateUserArgs
	}

	DeleteInvitationByCodeFn     func(ctx context.Context, tx db.Transaction, code string) error
	DeleteInvitationByCodeCalled struct {
		Count int
		With  DeleteInvitationByCodeArgs
	}

	GetInvitationByCodeFn     func(ctx context.Context, tx db.Transaction, code string) (*invitation.Invitation, error)
	GetInvitationByCodeCalled struct {
		Count int
		With  GetInvitationByCodeArgs
	}

	SaveInvitationFn     func(ctx context.Context, tx db.Transaction, i *invitation.Invitation, hashedCode string) error
	SaveInvitationCalled struct {
		Count int
		With  SaveInvitationArgs
	}
}

func (repo *Repository) AddMemberToOrganization(ctx context.Context, tx db.Transaction, userID, orgID string) error {
	if repo.AddMemberToOrganizationFn == nil {
		return mockMethodNotDefined("AddMemberToOrganization")
	}
	repo.AddMemberToOrganizationCalled.Count++
	repo.AddMemberToOrganizationCalled.With = AddMemberToOrganizationArgs{OrganizationID: orgID, UserID: userID}
	return repo.AddMemberToOrganizationFn(ctx, tx, userID, orgID)
}

func (repo *Repository) CreateAward(ctx context.Context, tx db.Transaction, a *award.Award) error {
	if repo.CreateAwardFn == nil {
		return mockMethodNotDefined("CreateAward")
	}
	repo.CreateAwardCalled.Count++
	repo.CreateAwardCalled.With = CreateAwardArgs{Award: a}
	return repo.CreateAwardFn(ctx, tx, a)
}

func (repo *Repository) CreateCheevo(ctx context.Context, tx db.Transaction, cheevo *cheevo.Cheevo) error {
	if repo.CreateCheevoFn == nil {
		return mockMethodNotDefined("CreateCheevo")
	}
	repo.CreateCheevoCalled.Count++
	repo.CreateCheevoCalled.With = CreateCheevoArgs{Cheevo: cheevo}
	return repo.CreateCheevoFn(ctx, tx, cheevo)
}

func (repo *Repository) CreateInvitation(ctx context.Context, tx db.Transaction, i *invitation.Invitation, hashedCode string) error {
	if repo.CreateInvitationFn == nil {
		return mockMethodNotDefined("CreateInvitation")
	}
	repo.CreateInvitationCalled.Count++
	repo.CreateInvitationCalled.With = CreateInvitationArgs{Invitation: i, HashedCode: hashedCode}
	return repo.CreateInvitationFn(ctx, tx, i, hashedCode)
}

func (repo *Repository) CreateMembership(ctx context.Context, tx db.Transaction, m *membership.Membership) error {
	if repo.CreateMembershipFn == nil {
		return mockMethodNotDefined("CreateMembership")
	}
	repo.CreateMembershipCalled.Count++
	repo.CreateMembershipCalled.With = CreateMembershipArgs{Membership: m}
	return repo.CreateMembershipFn(ctx, tx, m)
}

func (repo *Repository) CreateOrganization(ctx context.Context, tx db.Transaction, org *organization.Organization) error {
	if repo.CreateOrganizationFn == nil {
		return mockMethodNotDefined("CreateOrganization")
	}
	repo.CreateOrganizationCalled.Count++
	repo.CreateOrganizationCalled.With = CreateOrganizationArgs{Organization: org}
	return repo.CreateOrganizationFn(ctx, tx, org)
}

func (repo *Repository) CreateUser(ctx context.Context, tx db.Transaction, u *user.User, hashedPassword string) error {
	if repo.CreateUserFn == nil {
		return mockMethodNotDefined("CreateUser")
	}
	repo.CreateUserCalled.Count++
	repo.CreateUserCalled.With = CreateUserArgs{User: u, PasswordHash: hashedPassword}
	return repo.CreateUserFn(ctx, tx, u, hashedPassword)
}

func (repo *Repository) DeleteInvitationByCode(ctx context.Context, tx db.Transaction, code string) error {
	if repo.DeleteInvitationByCodeFn == nil {
		return mockMethodNotDefined("DeleteInvitationByCode")
	}
	repo.DeleteInvitationByCodeCalled.Count++
	repo.DeleteInvitationByCodeCalled.With = DeleteInvitationByCodeArgs{Code: code}
	return repo.DeleteInvitationByCodeFn(ctx, tx, code)
}

func (repo *Repository) GetInvitationByCode(ctx context.Context, tx db.Transaction, code string) (*invitation.Invitation, error) {
	if repo.GetInvitationByCodeFn == nil {
		return nil, mockMethodNotDefined("GetInvitationByCode")
	}
	repo.GetInvitationByCodeCalled.Count++
	repo.GetInvitationByCodeCalled.With = GetInvitationByCodeArgs{Code: code}
	return repo.GetInvitationByCodeFn(ctx, tx, code)
}

func (repo *Repository) SaveInvitation(ctx context.Context, tx db.Transaction, i *invitation.Invitation, hashedCode string) error {
	if repo.GetInvitationByCodeFn == nil {
		return mockMethodNotDefined("SaveInvitation")
	}
	repo.SaveInvitationCalled.Count++
	repo.SaveInvitationCalled.With = SaveInvitationArgs{Invitation: i, HashedCode: hashedCode}
	return repo.SaveInvitationFn(ctx, tx, i, hashedCode)
}

func mockMethodNotDefined(funcName string) error {
	return fmt.Errorf("mock method %s is not defined", funcName)
}
