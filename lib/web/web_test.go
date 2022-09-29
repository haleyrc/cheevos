package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleErrorRespondsWithTheDefaultCodeWhenErrorIsntCoded(t *testing.T) {
	w := httptest.NewRecorder()
	err := fmt.Errorf("oops")
	handleError(w, err)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected to get code %d, but got %d.", http.StatusInternalServerError, w.Code)
	}
}

func TestHandleErrorRespondsWithTheCorrectCodeWhenErrorIsCoded(t *testing.T) {
	w := httptest.NewRecorder()
	err := codedError(http.StatusTeapot)
	handleError(w, err)
	if w.Code != http.StatusTeapot {
		t.Errorf("Expected to get code %d, but got %d.", http.StatusTeapot, w.Code)
	}
}

type codedError int

func (ce codedError) Code() int     { return int(ce) }
func (ce codedError) Error() string { return "oops" }
