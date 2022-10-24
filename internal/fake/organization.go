package fake

import (
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/domain"
)

func Organization(ownerID string) *domain.Organization {
	return &domain.Organization{
		ID:      uuid.New(),
		Name:    uniqify("TestOrg"),
		OwnerID: ownerID,
	}
}
