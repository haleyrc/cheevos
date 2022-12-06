package password

func New(s string) Password {
	return Password{s: s}
}

type Password struct {
	s string
}
