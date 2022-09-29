package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
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
			if w.Code != tc.want {
				t.Errorf("Expected to get code %d, but got %d.", tc.want, w.Code)
			}
		})
	}
}

func TestHandleErrorReturnsTheExpectedMessage(t *testing.T) {
	testcases := map[string]struct {
		input error
		want  string
	}{
		"when the error isn't messaged": {input: fmt.Errorf("oops"), want: "Unexpected error."},
		"when the error is messaged":    {input: messagedError("hello"), want: "hello"},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			handleError(w, tc.input)
			got := parseErrorMessage(t, w)
			if got != tc.want {
				t.Errorf("Expected to get message %q, but got %q.", tc.want, got)
			}
		})
	}
}

func parseErrorMessage(t *testing.T, w *httptest.ResponseRecorder) string {
	var resp struct {
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatal(err)
	}
	return resp.Error.Message
}

type messagedError string

func (me messagedError) Message() string { return string(me) }
func (me messagedError) Error() string   { return "oops" }

type codedError int

func (ce codedError) Code() int     { return int(ce) }
func (ce codedError) Error() string { return "oops" }
