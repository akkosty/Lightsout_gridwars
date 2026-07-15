package api

import (
	"testing"
)

func TestResponse_Success(t *testing.T) {
	resp := Response{
		Success: true,
		Message: "Operation successful",
		Data:    map[string]string{"key": "value"},
	}

	if !resp.Success {
		t.Error("Expected Success to be true")
	}

	if resp.Message != "Operation successful" {
		t.Errorf("Expected message 'Operation successful', got '%s'", resp.Message)
	}
}

func TestResponse_Error(t *testing.T) {
	err := &Error{
		Code:    404,
		Message: "Not found",
		Details: "User not found",
	}

	resp := Response{
		Success: false,
		Error:   err,
	}

	if resp.Success {
		t.Error("Expected Success to be false")
	}

	if resp.Error.Code != 404 {
		t.Errorf("Expected error code 404, got %d", resp.Error.Code)
	}
}

func TestError_Error(t *testing.T) {
	err := &Error{
		Code:    400,
		Message: "Bad request",
		Details: "Invalid input data",
	}

	expected := "Bad request"
	if err.Message != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Message)
	}
}

func TestErrNotFound(t *testing.T) {
	err := NewNotFoundError("user", "12345")

	expected := "user with ID 12345 not found"
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestErrAlreadyExists(t *testing.T) {
	err := NewAlreadyExistsError("username", "testuser")

	expected := "username with value 'testuser' already exists"
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestErrInvalidInput(t *testing.T) {
	err := NewInvalidInputError("email", "invalid format")

	expected := "invalid input for field: email: invalid format"
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestErrInvalidInput_NoDetails(t *testing.T) {
	err := NewInvalidInputError("age", "")

	expected := "invalid input for field: age"
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}
