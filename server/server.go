package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/haleyrc/cheevos/auth"
	"github.com/haleyrc/cheevos/cheevos"
	"github.com/haleyrc/cheevos/lib/web"
	"github.com/haleyrc/cheevos/roster"
)

type AuthorizationService interface {
	CanAwardCheevo(ctx context.Context, awarderID, recipientID, cheevoID string) error
	CanCreateCheevo(ctx context.Context, userID, orgID string) error
	CanGetCheevo(ctx context.Context, userID, cheevoID string) error
	CanInviteUsersToOrganization(ctx context.Context, userID, orgID string) error
	CanRefreshInvitation(ctx context.Context, userID, invitationID string) error
}

type AuthenticationService interface {
	SignUp(ctx context.Context, username, password string) (*auth.User, error)
}

type CheevosService interface {
	AwardCheevoToUser(ctx context.Context, recipientID, cheevoID string) error
	CreateCheevo(ctx context.Context, name, description, orgID string) (*cheevos.Cheevo, error)
	GetCheevo(ctx context.Context, id string) (*cheevos.Cheevo, error)
}

type RosterService interface {
	AcceptInvitation(ctx context.Context, userID, code string) error
	CreateOrganization(ctx context.Context, name, ownerID string) (*roster.Organization, error)
	DeclineInvitation(ctx context.Context, code string) error
	InviteUserToOrganization(ctx context.Context, email, orgID string) (*roster.Invitation, error)
	RefreshInvitation(ctx context.Context, invitationID string) (*roster.Invitation, error)
}

type Server struct {
	Auth    AuthServer
	Cheevos CheevosServer
	Roster  RosterServer

	mux *http.ServeMux
}

func (srv *Server) Start(ctx context.Context) error {
	srv.registerRoutes()
	if err := http.ListenAndServe(":8080", srv.mux); err != nil {
		return fmt.Errorf("Server.Start: %w", err)
	}
	return nil
}

func (srv *Server) registerRoutes() {
	if srv.mux == nil {
		srv.mux = http.NewServeMux()
	}

	srv.mux.HandleFunc("/auth/SignUp", web.ResponseHandler(srv.Auth.SignUp))

	srv.mux.HandleFunc("/cheevos/AwardCheevo", web.ResponseHandler(srv.Cheevos.AwardCheevo))
	srv.mux.HandleFunc("/cheevos/CreateCheevo", web.ResponseHandler(srv.Cheevos.CreateCheevo))
	srv.mux.HandleFunc("/cheevos/GetCheevo", web.ResponseHandler(srv.Cheevos.GetCheevo))

	srv.mux.HandleFunc("/roster/AcceptInvitation", web.ResponseHandler(srv.Roster.AcceptInvitation))
	srv.mux.HandleFunc("/roster/CreateOrganization", web.ResponseHandler(srv.Roster.CreateOrganization))
	srv.mux.HandleFunc("/roster/DeclineInvitation", web.ResponseHandler(srv.Roster.DeclineInvitation))
	srv.mux.HandleFunc("/roster/InviteUserToOrganization", web.ResponseHandler(srv.Roster.InviteUserToOrganization))
	srv.mux.HandleFunc("/roster/RefreshInvitation", web.ResponseHandler(srv.Roster.RefreshInvitation))
}
