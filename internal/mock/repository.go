package mock

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos/auth"
	"github.com/haleyrc/cheevos/cheevos"
	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/roster"
)

var _ = &auth.Service{Repo: &Repository{}}
var _ = &cheevos.Service{Repo: &Repository{}}
var _ = &roster.Service{Repo: &Repository{}}

type CreateAwardArgs struct {
	Award *cheevos.Award
}

type CreateCheevoArgs struct {
	Cheevo *cheevos.Cheevo
}

type CreateInvitationArgs struct {
	Invitation *roster.Invitation
	HashedCode string
}

type CreateMembershipArgs struct {
	Membership *roster.Membership
}

type CreateOrganizationArgs struct {
	Organization *roster.Organization
}

type CreateUserArgs struct {
	User         *auth.User
	PasswordHash string
}

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

type GetMemberArgs struct {
	OrganizationID string
	UserID         string
}

type GetUserArgs struct {
	ID string
}

type SaveInvitationArgs struct {
	Invitation *roster.Invitation
	HashedCode string
}

type Repository struct {
	CreateAwardFn     func(ctx context.Context, tx db.Tx, a *cheevos.Award) error
	CreateAwardCalled struct {
		Count int
		With  CreateAwardArgs
	}

	CreateCheevoFn     func(ctx context.Context, tx db.Tx, cheevo *cheevos.Cheevo) error
	CreateCheevoCalled struct {
		Count int
		With  CreateCheevoArgs
	}

	CreateInvitationFn     func(ctx context.Context, tx db.Tx, i *roster.Invitation, hashedCode string) error
	CreateInvitationCalled struct {
		Count int
		With  CreateInvitationArgs
	}

	CreateMembershipFn     func(ctx context.Context, tx db.Tx, m *roster.Membership) error
	CreateMembershipCalled struct {
		Count int
		With  CreateMembershipArgs
	}

	CreateOrganizationFn     func(ctx context.Context, tx db.Tx, org *roster.Organization) error
	CreateOrganizationCalled struct {
		Count int
		With  CreateOrganizationArgs
	}

	CreateUserFn     func(ctx context.Context, tx db.Tx, u *auth.User, hashedPassword string) error
	CreateUserCalled struct {
		Count int
		With  CreateUserArgs
	}

	DeleteInvitationByCodeFn     func(ctx context.Context, tx db.Tx, code string) error
	DeleteInvitationByCodeCalled struct {
		Count int
		With  DeleteInvitationByCodeArgs
	}

	GetCheevoFn     func(ctx context.Context, tx db.Tx, cheevo *cheevos.Cheevo, id string) error
	GetCheevoCalled struct {
		Count int
		With  GetCheevoArgs
	}

	GetInvitationFn     func(ctx context.Context, tx db.Tx, i *roster.Invitation, id string) error
	GetInvitationCalled struct {
		Count int
		With  GetInvitationArgs
	}

	GetInvitationByCodeFn     func(ctx context.Context, tx db.Tx, i *roster.Invitation, code string) error
	GetInvitationByCodeCalled struct {
		Count int
		With  GetInvitationByCodeArgs
	}

	GetMemberFn     func(ctx context.Context, tx db.Tx, m *roster.Membership, orgID, userID string) error
	GetMemberCalled struct {
		Count int
		With  GetMemberArgs
	}

	GetUserFn     func(ctx context.Context, tx db.Tx, u *auth.User, id string) error
	GetUserCalled struct {
		Count int
		With  GetUserArgs
	}

	SaveInvitationFn     func(ctx context.Context, tx db.Tx, i *roster.Invitation, hashedCode string) error
	SaveInvitationCalled struct {
		Count int
		With  SaveInvitationArgs
	}
}

func (repo *Repository) CreateAward(ctx context.Context, tx db.Tx, a *cheevos.Award) error {
	if repo.CreateAwardFn == nil {
		return mockMethodNotDefined("CreateAward")
	}
	repo.CreateAwardCalled.Count++
	repo.CreateAwardCalled.With = CreateAwardArgs{Award: a}
	return repo.CreateAwardFn(ctx, tx, a)
}

func (repo *Repository) CreateCheevo(ctx context.Context, tx db.Tx, cheevo *cheevos.Cheevo) error {
	if repo.CreateCheevoFn == nil {
		return mockMethodNotDefined("CreateCheevo")
	}
	repo.CreateCheevoCalled.Count++
	repo.CreateCheevoCalled.With = CreateCheevoArgs{Cheevo: cheevo}
	return repo.CreateCheevoFn(ctx, tx, cheevo)
}

func (repo *Repository) CreateInvitation(ctx context.Context, tx db.Tx, i *roster.Invitation, hashedCode string) error {
	if repo.CreateInvitationFn == nil {
		return mockMethodNotDefined("CreateInvitation")
	}
	repo.CreateInvitationCalled.Count++
	repo.CreateInvitationCalled.With = CreateInvitationArgs{Invitation: i, HashedCode: hashedCode}
	return repo.CreateInvitationFn(ctx, tx, i, hashedCode)
}

func (repo *Repository) CreateMembership(ctx context.Context, tx db.Tx, m *roster.Membership) error {
	if repo.CreateMembershipFn == nil {
		return mockMethodNotDefined("CreateMembership")
	}
	repo.CreateMembershipCalled.Count++
	repo.CreateMembershipCalled.With = CreateMembershipArgs{Membership: m}
	return repo.CreateMembershipFn(ctx, tx, m)
}

func (repo *Repository) CreateOrganization(ctx context.Context, tx db.Tx, org *roster.Organization) error {
	if repo.CreateOrganizationFn == nil {
		return mockMethodNotDefined("CreateOrganization")
	}
	repo.CreateOrganizationCalled.Count++
	repo.CreateOrganizationCalled.With = CreateOrganizationArgs{Organization: org}
	return repo.CreateOrganizationFn(ctx, tx, org)
}

func (repo *Repository) CreateUser(ctx context.Context, tx db.Tx, u *auth.User, hashedPassword string) error {
	if repo.CreateUserFn == nil {
		return mockMethodNotDefined("CreateUser")
	}
	repo.CreateUserCalled.Count++
	repo.CreateUserCalled.With = CreateUserArgs{User: u, PasswordHash: hashedPassword}
	return repo.CreateUserFn(ctx, tx, u, hashedPassword)
}

func (repo *Repository) DeleteInvitationByCode(ctx context.Context, tx db.Tx, code string) error {
	if repo.DeleteInvitationByCodeFn == nil {
		return mockMethodNotDefined("DeleteInvitationByCode")
	}
	repo.DeleteInvitationByCodeCalled.Count++
	repo.DeleteInvitationByCodeCalled.With = DeleteInvitationByCodeArgs{Code: code}
	return repo.DeleteInvitationByCodeFn(ctx, tx, code)
}

func (repo *Repository) GetCheevo(ctx context.Context, tx db.Tx, cheevo *cheevos.Cheevo, id string) error {
	if repo.GetCheevoFn == nil {
		return mockMethodNotDefined("GetCheevo")
	}
	repo.GetCheevoCalled.Count++
	repo.GetCheevoCalled.With = GetCheevoArgs{ID: id}
	return repo.GetCheevoFn(ctx, tx, cheevo, id)
}

func (repo *Repository) GetInvitation(ctx context.Context, tx db.Tx, i *roster.Invitation, id string) error {
	if repo.GetInvitationFn == nil {
		return mockMethodNotDefined("GetInvitation")
	}
	repo.GetInvitationCalled.Count++
	repo.GetInvitationCalled.With = GetInvitationArgs{ID: id}
	return repo.GetInvitationFn(ctx, tx, i, id)
}

func (repo *Repository) GetInvitationByCode(ctx context.Context, tx db.Tx, i *roster.Invitation, code string) error {
	if repo.GetInvitationByCodeFn == nil {
		return mockMethodNotDefined("GetInvitationByCode")
	}
	repo.GetInvitationByCodeCalled.Count++
	repo.GetInvitationByCodeCalled.With = GetInvitationByCodeArgs{Code: code}
	return repo.GetInvitationByCodeFn(ctx, tx, i, code)
}

func (repo *Repository) GetMember(ctx context.Context, tx db.Tx, m *roster.Membership, orgID, userID string) error {
	if repo.GetMemberFn == nil {
		return mockMethodNotDefined("GetMember")
	}
	repo.GetMemberCalled.Count++
	repo.GetMemberCalled.With = GetMemberArgs{OrganizationID: orgID, UserID: userID}
	return repo.GetMemberFn(ctx, tx, m, orgID, userID)
}

func (repo *Repository) GetUser(ctx context.Context, tx db.Tx, u *auth.User, id string) error {
	if repo.GetUserFn == nil {
		return mockMethodNotDefined("GetUser")
	}
	repo.GetUserCalled.Count++
	repo.GetUserCalled.With = GetUserArgs{ID: id}
	return repo.GetUserFn(ctx, tx, u, id)
}

func (repo *Repository) SaveInvitation(ctx context.Context, tx db.Tx, i *roster.Invitation, hashedCode string) error {
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
