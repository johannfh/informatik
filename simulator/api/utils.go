package api

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	StatusCode int `json:"statusCode"`
	Msg        any `json:"msg"`
}

func (err ApiError) Error() string {
	return fmt.Sprintf("api error: %d", err.StatusCode)
}

func NewApiError(statusCode int, err error) ApiError {
	return ApiError{
		StatusCode: statusCode,
		Msg:        err.Error(),
	}
}

func MakeHandlerFunc(hf HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := hf(w, r)
		if err == nil {
			return
		}

	}
}
