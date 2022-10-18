package fake

import (
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/lib/hash"
	"github.com/haleyrc/cheevos/internal/lib/random"
)

func User() *cheevos.User {
	return &cheevos.User{
		ID:       uuid.New(),
		Username: uniqify("TestUser"),
	}
}

func Password() (string, string) {
	password := random.String(8)
	hash := hash.Generate(password)
	return password, hash
}
