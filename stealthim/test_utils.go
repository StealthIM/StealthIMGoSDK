package stealthim

import (
	"crypto/rand"
	"encoding/hex"
)

// generateRandomUsername generates a random username for testing
func generateRandomUsername() string {
	bytes := make([]byte, 8)
	if _, err := rand.Read(bytes); err != nil {
		return "testuser123"
	}
	// Ensure username is at least 3 characters and at most 20 characters
	// "testuser_" is 9 characters, hex.EncodeToString(bytes) generates 16 characters
	// Total will be 25 characters, so we truncate to meet the 20 character limit
	username := "testuser_" + hex.EncodeToString(bytes)
	if len(username) > 20 {
		return username[:20]
	}
	return username
}