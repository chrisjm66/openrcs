package service

// For service bindings
type ApiError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

type ErrorCode string

const (
	LAYOUT_NOT_FOUND  ErrorCode = "LAYOUT_NOT_FOUND"
	RUNTIME_NOT_FOUND ErrorCode = "RUNTIME_NOT_FOUND"
)
