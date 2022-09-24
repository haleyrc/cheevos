package web

import (
	"encoding/json"
	"net/http"
)

type Response interface{}

type ServerFunc func(w http.ResponseWriter, r *http.Request) (Response, error)

func ResponseHandler(f ServerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := f(w, r)
		if err != nil {
			handleError(w, err)
			return
		}
		handleResponse(w, resp)
	}
}

type webError struct {
	Message string `json:"message"`
}

type errorResponse struct {
	Error webError `json:"error"`
}

func handleError(w http.ResponseWriter, err error) {
	code := errorCode(err)
	respondWithJSON(w, code, errorResponse{
		Error: webError{Message: err.Error()},
	})
}

type successResponse struct {
	Data interface{} `json:"data"`
}

func handleResponse(w http.ResponseWriter, data interface{}) {
	respondWithJSON(w, http.StatusOK, successResponse{
		Data: data,
	})
}

type Coded interface {
	Code() int
}

func errorCode(err error) int {
	code := http.StatusInternalServerError
	if err, ok := err.(Coded); ok {
		code = err.Code()
	}
	return code
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
