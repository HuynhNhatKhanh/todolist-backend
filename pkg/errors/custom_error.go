package errors

import "errors"

var (
	ErrInvalidRequest = errors.New("invalid request data")
	ErrTaskNotFound   = errors.New("task not found")
	ErrInvalidTaskID  = errors.New("invalid task ID")
	ErrTitleRequired  = errors.New("title is required")
	ErrInvalidJSON    = errors.New("invalid JSON format")
	ErrDatabase       = errors.New("database error")
)
