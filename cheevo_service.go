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

// AwardCheevoToUserRequest represents the parameters for awarding a specific
// Cheevo to a User.
type AwardCheevoToUserRequest struct {
	// The ID of the Cheevo to be awarded.
	Cheevo string

	// The ID of the User to receive the Cheevo.
	Awardee string

	// The ID of the User awarding the Cheevo.
	Awarder string
}

func (req *AwardCheevoToUserRequest) normalize() {
	req.Cheevo = strings.TrimSpace(req.Cheevo)
	req.Awardee = strings.TrimSpace(req.Awardee)
	req.Awarder = strings.TrimSpace(req.Awarder)
}

func (req *AwardCheevoToUserRequest) validate() error {
	if req.Cheevo == "" {
		return fmt.Errorf("invalid: cheevo is blank")
	}

	if req.Awardee == "" {
		return fmt.Errorf("invalid: awardee is blank")
	}

	if req.Awarder == "" {
		return fmt.Errorf("invalid: awarder is blank")
	}

	return nil
}

// AwardCheevoToUserResponse is returned when a Cheevo is successfully awarded
// to a User
type AwardCheevoToUserResponse struct {
	// The complete Cheevo that was awarded. The Cheevo's statistics reflect the
	// latest values after the award.
	Cheevo *Cheevo

	// The complete User that received the Cheevo. The User's statistics reflect
	// the latest values after the award.
	User *User
}

// AwardCheevoToUser awards a specific Cheevo to a User. Statistics for this
// event are bidirectional; a Cheevo "tracks" the number of Users that have
// received it and Users "track" how many Cheevos they have received.
func (cs *CheevoService) AwardCheevoToUser(ctx context.Context, req AwardCheevoToUserRequest) (*AwardCheevoToUserResponse, error) {
	req.normalize()

	if err := req.validate(); err != nil {
		return nil, fmt.Errorf("create cheevo failed: %w", err)
	}

	var cheevo *Cheevo
	var awardee *User
	err := cs.DB.Call(ctx, func(ctx context.Context, tx Transaction) error {
		var err error

		cheevo, err = tx.GetCheevo(ctx, req.Cheevo)
		if err != nil {
			return err
		}

		awardee, err = tx.GetUser(ctx, req.Awardee)
		if err != nil {
			return err
		}

		awarder, err := tx.GetUser(ctx, req.Awarder)
		if err != nil {
			return err
		}

		if err = tx.AwardCheevoToUser(ctx, cheevo, awardee, awarder); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("create organization failed: %w", err)
	}

	resp := &AwardCheevoToUserResponse{
		Cheevo: cheevo,
		User:   awardee,
	}
	return resp, nil
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
