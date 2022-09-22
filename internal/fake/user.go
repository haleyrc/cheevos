package fake

import (
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/auth"
	"github.com/haleyrc/cheevos/lib/hash"
	"github.com/haleyrc/cheevos/lib/random"
)

func User() (*auth.User, string, string) {
	password := random.String(8)
	hash := hash.Generate(password)
	u := &auth.User{
		ID:       uuid.New(),
		Username: uniqify("TestUser"),
	}
	return u, password, hash
}
