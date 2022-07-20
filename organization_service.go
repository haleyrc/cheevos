package cheevos

import (
	"context"
	"fmt"
	"strings"

	"github.com/pborman/uuid"
)

// OrganizationService represents the main entrypoint for managing
// organizations.
type OrganizationService struct {
	DB Database
}

// CreateOrganizationRequest represents the parameters for creating a new
// organization.
type CreateOrganizationRequest struct {
	// The name of the organization does not have to be unique, but should be
	// representative of the group being created.
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

// CreateOrganizationResponse is returned when an organization is successfully
// created.
type CreateOrganizationResponse struct {
	// The complete persisted organization. The ID returned on the model is a
	// unique identifer for the organization for use in further operations.
	Organization *Organization
}

// CreateOrganization creates a new organization and persists it to the
// database. It returns a response containing the new organization if
// successful.
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
