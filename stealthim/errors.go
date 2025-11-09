package stealthim

import (
	"errors"
	"fmt"
)

// StealthError represents a StealthIM API error
type StealthError struct {
	Code int
	Msg  string
}

func (e *StealthError) Error() string {
	return fmt.Sprintf("StealthIM Error %d: %s", e.Code, e.Msg)
}

// IsSuccess checks if the result is successful
func (r *Result) IsSuccess() bool {
	return r.Code == 800
}

// ToError converts a Result to an error if not successful
func (r *Result) ToError() error {
	if r.IsSuccess() {
		return nil
	}
	return &StealthError{Code: r.Code, Msg: r.Msg}
}

// Common error codes
var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserPasswordError = errors.New("user password error")
	ErrPermissionDenied  = errors.New("permission denied")
	ErrGroupNotFound     = errors.New("group not found")
	ErrFileNotFound      = errors.New("file not found")
)