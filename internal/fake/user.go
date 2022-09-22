package fake

import (
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/auth"
	"github.com/haleyrc/cheevos/lib/hash"
	"github.com/haleyrc/cheevos/lib/random"
)

func User() *auth.User {
	return &auth.User{
		ID:       uuid.New(),
		Username: uniqify("TestUser"),
	}
}

func Password() (string, string) {
	password := random.String(8)
	hash := hash.Generate(password)
	return password, hash
}
