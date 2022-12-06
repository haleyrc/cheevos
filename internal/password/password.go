package password

import (
	"strings"

	"github.com/haleyrc/pkg/hash"
)

func New(s string) Password {
	s = strings.TrimSpace(s)
	return Password{s: s}
}

type Password struct {
	s string
}

func (p Password) Hash() string {
	return hash.Generate(p.s)
}
