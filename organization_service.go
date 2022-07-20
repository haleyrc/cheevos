package cheevos

import (
	"context"
	"fmt"
	"strings"

	"github.com/pborman/uuid"
)

type OrganizationService struct {
	DB Database
}

type CreateOrganizationRequest struct {
	Name string
}

func (req *CreateOrganizationRequest) normalize() {
	req.Name = strings.TrimSpace(req.Name)
}

func (req *CreateOrganizationRequest) validate() error {
	if req.Name == "" {
		return fmt.Errorf("invalid: name is blank")
	}

	return nil
}

type CreateOrganizationResponse struct {
	Organization *Organization
}

func (os *OrganizationService) CreateOrganization(ctx context.Context, req CreateOrganizationRequest) (*CreateOrganizationResponse, error) {
	req.normalize()

	if err := req.validate(); err != nil {
		return nil, fmt.Errorf("create organization failed: %w", err)
	}

	org := &Organization{
		ID:   uuid.New(),
		Name: req.Name,
	}
	err := os.DB.Call(ctx, func(ctx context.Context, tx Transaction) error {
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("create organization failed: %w", err)
	}

	resp := &CreateOrganizationResponse{Organization: org}
	return resp, nil
}
