package fake

import (
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/roster"
)

func Organization(ownerID string) *roster.Organization {
	return &roster.Organization{
		ID:      uuid.New(),
		Name:    uniqify("TestOrg"),
		OwnerID: ownerID,
	}
}
