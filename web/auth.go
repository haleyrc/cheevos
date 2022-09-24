package web

import (
	"net/http"

	"github.com/haleyrc/cheevos/auth"
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

func (as *AuthServer) SignUp(w http.ResponseWriter, r *http.Request) (Response, error) {
	ctx := r.Context()

	var req SignUpRequest
	if err := decodeJSON(&req, r.Body); err != nil {
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
