package domain

import (
	"github.com/haleyrc/pkg/hash"
	"github.com/haleyrc/pkg/random"
	"github.com/haleyrc/pkg/time"

	"github.com/haleyrc/cheevos/internal/lib/stringutil"
)

func NewInvitationCode() InvitationCode {
	return InvitationCode{
		Plaintext: random.String(InvitationCodeLength),
		Expires:   time.Now().Add(InvitationValidFor),
	}
}

type InvitationCode struct {
	Plaintext string
	Expires   time.Time
}

func (ic InvitationCode) Expired() bool {
	return ic.Expires.Before(time.Now())
}

func (ic InvitationCode) Hash() string {
	return hash.Generate(ic.Plaintext)
}

type Invitation struct {
	ID string

	Email string

	OrganizationID string

	Code InvitationCode
}

func (i *Invitation) Model() string { return "Invitation" }

func (i *Invitation) Normalize() {
	i.Email = stringutil.MakeSafe(i.Email)
}

func (i *Invitation) Refresh() {
	i.Code = NewInvitationCode()
}

func (i *Invitation) Validate() error {
	i.Normalize()

	ve := NewValidationError(i)

	if i.ID == "" {
		ve.Add("ID", "ID can't be blank.")
	}

	if i.Email == "" {
		ve.Add("Email", "Email can't be blank.")
	}

	if i.OrganizationID == "" {
		ve.Add("OrganizationID", "Organization ID can't be blank.")
	}

	// TODO: Right now we're just assuming codes are good because we only generate
	// and handle these internally. We could theoretically still do validation to
	// ensure data integrity, but I don't want to think about it any more at the
	// moment.
	// if i.Code.Expires.IsZero() {
	//   ve.Add("Expires", "Expiration time can't be blank.")
	// }

	return ve.Error()
}
