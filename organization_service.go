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

// AddUserToOrganizationRequest represents the parameters for adding a User to
// an Organization.
type AddUserToOrganizationRequest struct {
	// The ID of the Organization the User will be joining.
	Organization string

	// The ID of the User who is joining the Organization.
	User string
}

func (req *AddUserToOrganizationRequest) normalize() {
	req.Organization = strings.TrimSpace(req.Organization)
	req.User = strings.TrimSpace(req.User)
}

func (req *AddUserToOrganizationRequest) validate() error {
	if req.Organization == "" {
		return fmt.Errorf("invalid: organization is blank")
	}
	if req.User == "" {
		return fmt.Errorf("invalid: user is blank")
	}
	return nil
}

// AddUserToOrganizationResponse is returned when a User is successfully added
// to an Organization.
type AddUserToOrganizationResponse struct {
	// The complete Organization that received ther User. The Organization's
	// statistics reflect the latest values after the User has joined.
	Organization *Organization

	// The complete User that joined the Organization. The User's statistics
	// reflect the latest values after joining.
	User *User
}

// AddUserToOrganization adds a User to a specific Organization. Statistics for
// this event are bidirectional: an Organization "tracks" how many users it has
// and Users "track" how many Organizations they belong to.
func (os *OrganizationService) AddUserToOrganization(ctx context.Context, req AddUserToOrganizationRequest) (*AddUserToOrganizationResponse, error) {
	req.normalize()

	if err := req.validate(); err != nil {
		return nil, fmt.Errorf("add user to organization failed: %w", err)
	}

	var org *Organization
	var user *User
	err := os.DB.Call(ctx, func(ctx context.Context, tx Transaction) error {
		var err error

		if err = tx.AddUserToOrganization(ctx, req.Organization, req.User); err != nil {
			return err
		}

		org, err = tx.GetOrganization(ctx, req.Organization)
		if err != nil {
			return err
		}

		user, err = tx.GetUser(ctx, req.User)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("add user to organization failed: %w", err)
	}

	resp := &AddUserToOrganizationResponse{
		Organization: org,
		User:         user,
	}
	return resp, nil
}

// CreateOrganizationRequest represents the parameters for creating a new
// organization.
type CreateOrganizationRequest struct {
	// The name of the organization does not have to be unique, but should be
	// representative of the group being created.
	Name string

	// The User that owns this organization by virtue of being the person that
	// created it.
	Owner string
}

// We don't have to normalize Owner here for the same reason we don't have to
// test for a blank one: this isn't supplied by a user so it will either exist
// or it won't.
func (req *CreateOrganizationRequest) normalize() {
	req.Name = strings.TrimSpace(req.Name)
}

func (req *CreateOrganizationRequest) validate() error {
	if req.Name == "" {
		return fmt.Errorf("invalid: name is blank")
	}

	if req.Owner == "" {
		return fmt.Errorf("invalid: owner is blank")
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
		ID:    uuid.New(),
		Name:  req.Name,
		Owner: req.Owner,
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
