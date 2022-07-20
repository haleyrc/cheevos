package cheevos

import (
	"context"
	"fmt"
	"strings"

	"github.com/pborman/uuid"
)

type CheevoService struct {
	DB Database
}

type CreateCheevoRequest struct {
	Name        string
	Description string
}

func (req *CreateCheevoRequest) normalize() {
	req.Name = strings.TrimSpace(req.Name)
	req.Description = strings.TrimSpace(req.Description)
}

func (req CreateCheevoRequest) validate() error {
	if req.Name == "" {
		return fmt.Errorf("invalid: name is blank")
	}

	if req.Description == "" {
		return fmt.Errorf("invalid: description is blank")
	}

	return nil
}

type CreateCheevoResponse struct {
	Cheevo *Cheevo
}

func (cs *CheevoService) CreateCheevo(ctx context.Context, req CreateCheevoRequest) (*CreateCheevoResponse, error) {
	req.normalize()

	if err := req.validate(); err != nil {
		return nil, fmt.Errorf("create cheevo failed: %w", err)
	}

	cheevo := &Cheevo{
		ID:          uuid.New(),
		Name:        req.Name,
		Description: req.Description,
	}
	err := cs.DB.Call(ctx, func(ctx context.Context, tx Transaction) error {
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("create organization failed: %w", err)
	}

	resp := &CreateCheevoResponse{Cheevo: cheevo}
	return resp, nil
}
