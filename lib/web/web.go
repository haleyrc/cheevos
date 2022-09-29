package web

import (
	"encoding/json"
	"net/http"
)

const DefaultErrorMessage = "Unexpected error."

type Messaged interface {
	Message() string
}

type Coded interface {
	Code() int
}

type Data interface{}

type Error struct {
	Message string `json:"message"`
}

type Response struct {
	Error *Error      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

type ServerFunc func(w http.ResponseWriter, r *http.Request) (Data, error)

func ResponseHandler(f ServerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := f(w, r)
		if err != nil {
			handleError(w, err)
			return
		}
		handleData(w, data)
	}
}

func errorCode(err error) int {
	code := http.StatusInternalServerError
	if err, ok := err.(Coded); ok {
		code = err.Code()
	}
	return code
}

func errorMessage(err error) string {
	msg := DefaultErrorMessage
	if err, ok := err.(Messaged); ok {
		msg = err.Message()
	}
	return msg
}

func handleError(w http.ResponseWriter, err error) {
	code := errorCode(err)
	message := errorMessage(err)
	respondWithJSON(w, code, Response{
		Error: &Error{Message: message},
	})
}

func handleData(w http.ResponseWriter, data interface{}) {
	respondWithJSON(w, http.StatusOK, Response{Data: data})
}

func respondWithJSON(w http.ResponseWriter, code int, body interface{}) {
	bytes, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write(bytes); err != nil {
		panic(err)
	}
}
