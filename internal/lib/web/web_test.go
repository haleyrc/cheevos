package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/haleyrc/cheevos/internal/assert"
)

func TestHandleErrorReturnsTheExpectedCode(t *testing.T) {
	testcases := map[string]struct {
		input error
		want  int
	}{
		"when the error isn't coded": {input: fmt.Errorf("oops"), want: 500},
		"when the error is coded":    {input: codedError(http.StatusTeapot), want: http.StatusTeapot},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			handleError(w, tc.input)
			assert.Int(t, "code", w.Code).Equals(tc.want)
		})
	}
}

func TestHandleErrorReturnsTheExpectedMessage(t *testing.T) {
	testcases := map[string]struct {
		input error
		want  string
	}{
		"when the error isn't messaged": {input: fmt.Errorf("oops"), want: "Something went wrong."},
		"when the error is messaged":    {input: messagedError("hello"), want: "hello"},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			handleError(w, tc.input)
			got := parseErrorMessage(t, w)
			assert.String(t, "message", got).Equals(tc.want)
		})
	}
}

func parseErrorMessage(t *testing.T, w *httptest.ResponseRecorder) string {
	t.Helper()
	var resp struct {
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	body, err := io.ReadAll(w.Body)
	if err != nil {
		assert.Error(t, err).IsUnexpected()
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		assert.Error(t, err).IsUnexpected()
	}
	return resp.Error.Message
}

type messagedError string

func (me messagedError) Message() string { return string(me) }
func (me messagedError) Error() string   { return "oops" }

type codedError int

func (ce codedError) Code() int     { return int(ce) }
func (ce codedError) Error() string { return "oops" }
