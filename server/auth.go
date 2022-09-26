package server

import (
	"net/http"

	"github.com/haleyrc/cheevos/auth"
	"github.com/haleyrc/cheevos/lib/json"
	"github.com/haleyrc/cheevos/lib/web"
)

type AuthServer struct {
	Auth auth.Service
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
		return nil, err
	}

	user, err := as.Auth.SignUp(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	resp := SignUpResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	return resp, nil
}
