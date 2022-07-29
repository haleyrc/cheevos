package fake

import (
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos"
)

type OrganizationOption func(*cheevos.Organization)

func WithOwner(user *cheevos.User) OrganizationOption {
	return func(org *cheevos.Organization) { org.Owner = user.ID }
}

func Organization(opts ...OrganizationOption) *cheevos.Organization {
	org := &cheevos.Organization{ID: uuid.New(), Name: randomWord("Org")}
	for _, opt := range opts {
		opt(org)
	}
	return org
}

type CheevoOption func(*cheevos.Cheevo)

func WithOrganization(org *cheevos.Organization) CheevoOption {
	return func(cheevo *cheevos.Cheevo) { cheevo.Organization = org.ID }
}

func Cheevo(opts ...CheevoOption) *cheevos.Cheevo {
	cheevo := &cheevos.Cheevo{
		ID:          uuid.New(),
		Name:        randomWord("Cheevo"),
		Description: randomSentence(5),
	}
	for _, opt := range opts {
		opt(cheevo)
	}
	return cheevo
}
