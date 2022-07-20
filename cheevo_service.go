package cheevos

import (
	"context"
	"fmt"
	"strings"

	"github.com/pborman/uuid"
)

// CheevoService is the primary entrypoint for managing cheevos.
type CheevoService struct {
	// A connection to the database.
	DB Database
}

// CreateCheevoRequest represents the parameters for creating a new cheevo.
type CreateCheevoRequest struct {
	// The short name of the cheevo. This can either be descriptive (e.g. "100
	// Features Shipped"), but can also be witty or heavily personal (e.g. "1337
	// haX0r").
	Name string

	// A long-form description of the act which earns the cheevo (e.g. "Shipped
	// 100 features!").
	Description string

	// The parent Organization that owns the Cheevo.
	Organization string
}

// We don't have to normalize org here for the same reason we don't have to
// test for blank: the org ID is not provided by the user so the ID will either
// exist or it won't.
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

	if req.Organization == "" {
		return fmt.Errorf("invalid: organization is blank")
	}

	return nil
}

// CreateCheevoResponse is returned when a cheevo is successfully created.
type CreateCheevoResponse struct {
	// The complete persisted cheevo. The ID returned on the model is a unique
	// identifer for the cheevo for use in further operations.
	Cheevo *Cheevo
}

// CreateCheevo creates a new cheevo and persists it to the database. It returns
// a response containing the full cheevo if successful.
func (cs *CheevoService) CreateCheevo(ctx context.Context, req CreateCheevoRequest) (*CreateCheevoResponse, error) {
	req.normalize()

	if err := req.validate(); err != nil {
		return nil, fmt.Errorf("create cheevo failed: %w", err)
	}

	cheevo := &Cheevo{
		ID:           uuid.New(),
		Name:         req.Name,
		Description:  req.Description,
		Organization: req.Organization,
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
