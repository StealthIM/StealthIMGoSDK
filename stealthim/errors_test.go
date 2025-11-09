package stealthim

import (
	"testing"
)

// TestStealthErrorError tests the StealthError.Error method
func TestStealthErrorError(t *testing.T) {
	err := &StealthError{Code: 900, Msg: "server error"}
	expected := "StealthIM Error 900: server error"
	if err.Error() != expected {
		t.Errorf("Expected %s, got %s", expected, err.Error())
	}
}

// TestCommonErrors tests the common error variables
func TestCommonErrors(t *testing.T) {
	if ErrUserNotFound.Error() != "user not found" {
		t.Error("ErrUserNotFound has incorrect message")
	}
	
	if ErrUserAlreadyExists.Error() != "user already exists" {
		t.Error("ErrUserAlreadyExists has incorrect message")
	}
	
	if ErrUserPasswordError.Error() != "user password error" {
		t.Error("ErrUserPasswordError has incorrect message")
	}
	
	if ErrPermissionDenied.Error() != "permission denied" {
		t.Error("ErrPermissionDenied has incorrect message")
	}
	
	if ErrGroupNotFound.Error() != "group not found" {
		t.Error("ErrGroupNotFound has incorrect message")
	}
	
	if ErrFileNotFound.Error() != "file not found" {
		t.Error("ErrFileNotFound has incorrect message")
	}
}