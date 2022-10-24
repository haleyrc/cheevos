package fake

import (
	"github.com/haleyrc/pkg/hash"
	"github.com/haleyrc/pkg/random"
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/domain"
)

func User() *domain.User {
	return &domain.User{
		ID:       uuid.New(),
		Username: uniqify("TestUser"),
	}
}

func Password() (string, string) {
	password := random.String(8)
	hash := hash.Generate(password)
	return password, hash
}
