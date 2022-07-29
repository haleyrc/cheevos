package fake

import (
	"fmt"
	"time"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos"
)

func Username() string {
	return fmt.Sprintf("User%d", time.Now().UnixNano())
}

func User() *cheevos.User {
	return &cheevos.User{ID: uuid.New(), Username: randomWord("User")}
}
