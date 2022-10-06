package fake

import (
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/internal/lib/hash"
	"github.com/haleyrc/cheevos/internal/lib/random"
	"github.com/haleyrc/cheevos/internal/service"
)

func User() *service.User {
	return &service.User{
		ID:       uuid.New(),
		Username: uniqify("TestUser"),
	}
}

func Password() (string, string) {
	password := random.String(8)
	hash := hash.Generate(password)
	return password, hash
}
