package api

import "fmt"

// ErrNotFound represents a resource not found error
type ErrNotFound struct {
	Resource string
	ID       string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("%s with ID %s not found", e.Resource, e.ID)
}

// ErrAlreadyExists represents an already exists error
type ErrAlreadyExists struct {
	Resource string
	Value    string
}

func (e *ErrAlreadyExists) Error() string {
	return fmt.Sprintf("%s with value '%s' already exists", e.Resource, e.Value)
}

// ErrInvalidInput represents an invalid input error
type ErrInvalidInput struct {
	Field   string
	Details string
}

func (e *ErrInvalidInput) Error() string {
	msg := fmt.Sprintf("invalid input for field: %s", e.Field)
	if e.Details != "" {
		msg += ": " + e.Details
	}
	return msg
}

// NewNotFoundError creates a new ErrNotFound
func NewNotFoundError(resource, id string) error {
	return &ErrNotFound{Resource: resource, ID: id}
}

// NewAlreadyExistsError creates a new ErrAlreadyExists
func NewAlreadyExistsError(resource, value string) error {
	return &ErrAlreadyExists{Resource: resource, Value: value}
}

// NewInvalidInputError creates a new ErrInvalidInput
func NewInvalidInputError(field, details string) error {
	return &ErrInvalidInput{Field: field, Details: details}
}
