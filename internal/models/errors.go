package models

import (
	"net/http"
)

// ApiError wrapped httpStatusCode and error message
type ApiError struct {
	Msg        string `json:"error"`
	StatusCode int    `json:"-"`
}

func NewApiError(s string, statusCode int) *ApiError {
	return &ApiError{
		Msg:        s,
		StatusCode: statusCode,
	}
}

func (a *ApiError) Error() string {
	return a.Msg
}

var (
	ErrEmptyFileName        = NewApiError("file name must be set", http.StatusBadRequest)
	ErrNoData               = NewApiError("file data must be set", http.StatusBadRequest)
	ErrNotFound             = NewApiError("not found", http.StatusNotFound)
	ErrInternalServer       = NewApiError("internal server error", http.StatusInternalServerError)
	ErrInvalidMultipartForm = NewApiError("invalid multipart form", http.StatusBadRequest)
	ErrTooManyFiles         = NewApiError("too many files, while uploading", http.StatusBadRequest)
	ErrCouldNotOpenFile     = NewApiError("could not open file", http.StatusBadRequest)
	ErrCouldNotReadFile     = NewApiError("could not read file", http.StatusBadRequest)
)
