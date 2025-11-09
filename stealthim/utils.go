package stealthim

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// GenerateRandomString generates a random string of specified length
func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length/2+length%2) // Convert to required bytes (hex encoding doubles the length)
	if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
		return "", fmt.Errorf("failed to generate random string: %w", err)
	}
	return hex.EncodeToString(bytes)[:length], nil // Trim to exact length if needed
}

// ReadFile reads the entire content of a file
func ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

// WriteFile writes data to a file
func WriteFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}

// FileExists checks if a file exists
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}