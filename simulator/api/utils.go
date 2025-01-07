package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type ApiFunc func(http.ResponseWriter, *http.Request) error

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

func InvalidRequestData(errors map[string]string) ApiError {
	return ApiError{
		StatusCode: http.StatusUnprocessableEntity,
		Msg:        errors,
	}
}

func InvalidJSON() ApiError {
	return NewApiError(http.StatusUnprocessableEntity, fmt.Errorf("invalid JSON request data"))
}

func Make(hf ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := hf(w, r); err != nil {
			if apiErr, ok := err.(ApiError); ok {
				writeJSON(w, apiErr.StatusCode, apiErr)
			} else {
				statusCode := http.StatusInternalServerError
				errResp := map[string]any{
					"statusCode": statusCode,
					"msg":        http.StatusText(statusCode),
				}
				writeJSON(w, statusCode, errResp)
			}
			slog.Error("HTTP API error", "err", err, "path", r.URL.Path)
		}
	}
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) error {
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}
