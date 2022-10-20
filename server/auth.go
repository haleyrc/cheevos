package server

import (
	"net/http"

	"github.com/haleyrc/pkg/json"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/core"
	"github.com/haleyrc/cheevos/internal/lib/web"
)

type AuthServer struct {
	Auth cheevos.AuthService
}

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (as *AuthServer) SignUp(w http.ResponseWriter, r *http.Request) (web.Data, error) {
	ctx := r.Context()

	var req SignUpRequest
	if err := json.Decode(&req, r.Body); err != nil {
		return nil, core.NewBadRequestError(err)
	}

	user, err := as.Auth.SignUp(ctx, req.Username, req.Password)
	if err != nil {
		return nil, core.WrapError(err)
	}

	resp := SignUpResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	return resp, nil
}
