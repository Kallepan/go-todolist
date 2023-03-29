package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	StatusText     string `json:"status_text"`
	Message        string `json:"message"`
}

var (
	ErrMethodNotAllowed = &ErrorResponse{HTTPStatusCode: 405, StatusText: "Method Not Allowed", Message: "Method not allowed"}
	ErrNotFound         = &ErrorResponse{HTTPStatusCode: 404, StatusText: "Not Found", Message: "Resource not found"}
	ErrBadRequest       = &ErrorResponse{HTTPStatusCode: 400, StatusText: "Bad Request", Message: "Bad request"}
	ErrInternalServer   = &ErrorResponse{HTTPStatusCode: 500, StatusText: "Internal Server Error", Message: "Internal server error"}
)

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Bad Request",
		Message:        err.Error(),
	}
}

func ServerErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal Server Error",
		Message:        err.Error(),
	}
}
