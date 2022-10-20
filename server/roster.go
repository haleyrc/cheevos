package server

import (
	"net/http"

	"github.com/haleyrc/pkg/json"
	"github.com/haleyrc/pkg/time"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/core"
	"github.com/haleyrc/cheevos/internal/lib/web"
)

type RosterServer struct {
	Authz  AuthorizationService
	Roster cheevos.RosterService
}

type AcceptInvitationRequest struct {
	Code string `json:"code"`
}

type AcceptInvitationResponse struct{}

func (rs *RosterServer) AcceptInvitation(w http.ResponseWriter, r *http.Request) (web.Data, error) {
	ctx := r.Context()
	currentUser := GetCurrentUser(ctx)

	var req AcceptInvitationRequest
	if err := json.Decode(&req, r.Body); err != nil {
		return nil, core.NewBadRequestError(err)
	}

	if err := rs.Roster.AcceptInvitation(ctx, currentUser, req.Code); err != nil {
		return nil, core.WrapError(err)
	}

	return AcceptInvitationResponse{}, nil
}

type CreateOrganizationRequest struct {
	Name string `json:"name"`
}

type CreateOrganizationResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (rs *RosterServer) CreateOrganization(w http.ResponseWriter, r *http.Request) (web.Data, error) {
	ctx := r.Context()
	currentUser := GetCurrentUser(ctx)

	var req CreateOrganizationRequest
	if err := json.Decode(&req, r.Body); err != nil {
		return nil, core.NewBadRequestError(err)
	}

	org, err := rs.Roster.CreateOrganization(ctx, req.Name, currentUser)
	if err != nil {
		return nil, core.WrapError(err)
	}

	resp := CreateOrganizationResponse{
		ID:   org.ID,
		Name: org.Name,
	}

	return resp, nil
}

type DeclineInvitationRequest struct {
	Code string `json:"code"`
}

type DeclineInvitationResponse struct{}

func (rs *RosterServer) DeclineInvitation(w http.ResponseWriter, r *http.Request) (web.Data, error) {
	ctx := r.Context()

	var req DeclineInvitationRequest
	if err := json.Decode(&req, r.Body); err != nil {
		return nil, core.NewBadRequestError(err)
	}

	if err := rs.Roster.DeclineInvitation(ctx, req.Code); err != nil {
		return nil, core.WrapError(err)
	}

	return DeclineInvitationResponse{}, nil
}

type InviteUserToOrganizationRequest struct {
	Email          string `json:"email"`
	OrganizationID string `json:"organizationID"`
}

type InviteUserToOrganizationResponse struct {
	Expires time.Time `json:"expires"`
}

func (rs *RosterServer) InviteUserToOrganization(w http.ResponseWriter, r *http.Request) (web.Data, error) {
	ctx := r.Context()
	currentUser := GetCurrentUser(ctx)

	var req InviteUserToOrganizationRequest
	if err := json.Decode(&req, r.Body); err != nil {
		return nil, core.NewBadRequestError(err)
	}

	if err := rs.Authz.CanInviteUsersToOrganization(ctx, currentUser, req.OrganizationID); err != nil {
		return nil, core.WrapError(err)
	}

	invitation, err := rs.Roster.InviteUserToOrganization(ctx, req.Email, req.OrganizationID)
	if err != nil {
		return nil, core.WrapError(err)
	}

	resp := InviteUserToOrganizationResponse{
		Expires: invitation.Expires,
	}

	return resp, nil
}

type RefreshInvitationRequest struct {
	InvitationID string `json:"invitationID"`
}

type RefreshInvitationResponse struct {
	ID      string    `json:"id"`
	Expires time.Time `json:"expires"`
}

func (rs *RosterServer) RefreshInvitation(w http.ResponseWriter, r *http.Request) (web.Data, error) {
	ctx := r.Context()
	currentUser := GetCurrentUser(ctx)

	var req RefreshInvitationRequest
	if err := json.Decode(&req, r.Body); err != nil {
		return nil, core.NewBadRequestError(err)
	}

	if err := rs.Authz.CanRefreshInvitation(ctx, currentUser, req.InvitationID); err != nil {
		return nil, core.WrapError(err)
	}

	invitation, err := rs.Roster.RefreshInvitation(ctx, req.InvitationID)
	if err != nil {
		return nil, core.WrapError(err)
	}

	resp := RefreshInvitationResponse{
		ID:      invitation.ID,
		Expires: invitation.Expires,
	}

	return resp, nil
}
