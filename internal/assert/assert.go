package assert

import (
	"strings"
	"testing"
)

func Bool(t *testing.T, label string, value bool) BoolSubject {
	t.Helper()
	return BoolSubject{
		t:     t,
		label: label,
		value: value,
	}
}

type BoolSubject struct {
	t     *testing.T
	label string
	value bool
}

func (bs BoolSubject) IsFalse() {
	bs.t.Helper()
	if bs.value {
		bs.t.Errorf("Expected %s to be false, but it was true.", bs.label)
	}
}

func (bs BoolSubject) IsTrue() {
	bs.t.Helper()
	if !bs.value {
		bs.t.Errorf("Expected %s to be true, but it was false.", bs.label)
	}
}

func Error(t *testing.T, value error) ErrorSubject {
	t.Helper()
	return ErrorSubject{
		t:     t,
		value: value,
	}
}

type ErrorSubject struct {
	t     *testing.T
	value error
}

func (es ErrorSubject) IsNil() {
	es.t.Helper()
	if es.value != nil {
		es.t.Errorf("Expected error to be nil, but got %v.", es.value)
	}
}

func (es ErrorSubject) IsNotNil() {
	es.t.Helper()
	if es.value == nil {
		es.t.Errorf("Expected error not to be nil, but it was.")
	}
}

func (es ErrorSubject) IsUnexpected() {
	es.t.Helper()
	if es.value != nil {
		es.t.Fatal("unexpected error:", es.value)
	}
}

func Int(t *testing.T, label string, value int) IntSubject {
	t.Helper()
	return IntSubject{
		t:     t,
		label: label,
		value: value,
	}
}

type IntSubject struct {
	t     *testing.T
	label string
	value int
}

func (is IntSubject) Equals(value int) {
	is.t.Helper()
	if is.value != value {
		is.t.Errorf("Expected %s to be %d, but got %d.", is.label, value, is.value)
	}
}

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

func (ss StringSubject) Contains(value string) {
	ss.t.Helper()
	if !strings.Contains(ss.value, value) {
		ss.t.Errorf("Expected %s to contain %q, but it didn't.", ss.label, value)
	}
}

func (ss StringSubject) Equals(value string) {
	ss.t.Helper()
	if ss.value != value {
		ss.t.Errorf("Expected %s to be %q, but got %q.", ss.label, value, ss.value)
	}
}

func (ss StringSubject) NotBlank() {
	ss.t.Helper()
	if ss.value == "" {
		ss.t.Errorf("Expected %s to be blank, but it was %q.", ss.label, ss.value)
	}
}

func (ss StringSubject) NotEquals(value string) {
	ss.t.Helper()
	if ss.value == value {
		ss.t.Errorf("Expected %s to not be %s, but it was.", ss.label, value)
	}
}

func New(t *testing.T) Harness {
	return Harness{t: t}
}

type Harness struct {
	t *testing.T
}

func (h Harness) Bool(label string, value bool) BoolSubject {
	h.t.Helper()
	return Bool(h.t, label, value)
}

func (h Harness) Error(value error) ErrorSubject {
	h.t.Helper()
	return Error(h.t, value)
}

func (h Harness) Int(label string, value int) IntSubject {
	h.t.Helper()
	return Int(h.t, label, value)
}

func (h Harness) String(label, value string) StringSubject {
	h.t.Helper()
	return String(h.t, label, value)
}
