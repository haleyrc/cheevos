package assert

import (
	"testing"
)

func String(t *testing.T, label, value string) StringSubject {
	t.Helper()
	return StringSubject{
		t:     t,
		label: label,
		value: value,
	}
}

type StringSubject struct {
	t     *testing.T
	label string
	value string
}

func (ss StringSubject) Equals(value string) {
	ss.t.Helper()
	if ss.value != value {
		ss.t.Errorf("Expected %s to be %q, but got %q.", ss.label, value, ss.value)
	}
}

func New(t *testing.T) Harness {
	return Harness{t: t}
}

type Harness struct {
	t *testing.T
}

func (h Harness) String(label, value string) StringSubject {
	h.t.Helper()
	return String(h.t, label, value)
}
