package fake

import (
	"github.com/pborman/uuid"

	"github.com/haleyrc/pkg/hash"
	"github.com/haleyrc/pkg/random"

	"github.com/haleyrc/cheevos"
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
