package server

import (
	"net/http"

	"github.com/haleyrc/pkg/json"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/core"
	"github.com/haleyrc/cheevos/internal/lib/web"
)

type CheevosServer struct {
	Authz   AuthorizationService
	Cheevos cheevos.CheevosService
	Roster  cheevos.RosterService
}

type AwardCheevoRequest struct {
	CheevoID    string `json:"cheevoID"`
	RecipientID string `json:"recipientID"`
}

type AwardCheevoResponse struct {
	CheevoID    string `json:"cheevoID"`
	RecipientID string `json:"recipientID"`
}

func (cs *CheevosServer) AwardCheevo(w http.ResponseWriter, r *http.Request) (web.Data, error) {
	ctx := r.Context()
	currentUser := GetCurrentUser(ctx)

	var req AwardCheevoRequest
	if err := json.Decode(&req, r.Body); err != nil {
		return nil, core.NewBadRequestError(err)
	}

	if err := cs.Authz.CanAwardCheevo(ctx, currentUser, req.RecipientID, req.CheevoID); err != nil {
		return nil, core.WrapError(err)
	}

	if err := cs.Cheevos.AwardCheevoToUser(ctx, req.RecipientID, req.CheevoID); err != nil {
		return nil, core.WrapError(err)
	}

	resp := AwardCheevoResponse(req)
	return resp, nil
}

type CreateCheevoRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	OrganizationID string `json:"organizationID"`
}

type CreateCheevoResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	OrganizationID string `json:"organizationID"`
}

func (cs *CheevosServer) CreateCheevo(w http.ResponseWriter, r *http.Request) (web.Data, error) {
	ctx := r.Context()
	currentUser := GetCurrentUser(ctx)

	var req CreateCheevoRequest
	if err := json.Decode(&req, r.Body); err != nil {
		return nil, core.NewBadRequestError(err)
	}

	if err := cs.Authz.CanCreateCheevo(ctx, currentUser, req.OrganizationID); err != nil {
		return nil, core.WrapError(err)
	}

	cheevo, err := cs.Cheevos.CreateCheevo(ctx, req.Name, req.Description, req.OrganizationID)
	if err != nil {
		return nil, core.WrapError(err)
	}

	resp := CreateCheevoResponse{
		ID:             cheevo.ID,
		Name:           cheevo.Name,
		Description:    cheevo.Description,
		OrganizationID: cheevo.OrganizationID,
	}

	return resp, nil
}

type GetCheevoRequest struct {
	CheevoID string `json:"cheevoID"`
}

type GetCheevoResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	OrganizationID string `json:"organizationID"`
}

func (cs *CheevosServer) GetCheevo(w http.ResponseWriter, r *http.Request) (web.Data, error) {
	ctx := r.Context()
	currentUser := GetCurrentUser(ctx)

	var req GetCheevoRequest
	if err := json.Decode(&req, r.Body); err != nil {
		return nil, core.NewBadRequestError(err)
	}

	if err := cs.Authz.CanGetCheevo(ctx, currentUser, req.CheevoID); err != nil {
		return nil, core.WrapError(err)
	}

	cheevo, err := cs.Cheevos.GetCheevo(ctx, req.CheevoID)
	if err != nil {
		return nil, core.WrapError(err)
	}

	resp := GetCheevoResponse{
		ID:             cheevo.ID,
		Name:           cheevo.Name,
		Description:    cheevo.Description,
		OrganizationID: cheevo.OrganizationID,
	}

	return resp, nil
}
