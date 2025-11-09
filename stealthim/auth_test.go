package stealthim

import (
	"context"
	"os"
	"testing"
	"time"
)

// TestServerPing tests the Server.Ping method
func TestServerPing(t *testing.T) {
	// Get server URL from environment variable
	serverURL := os.Getenv("STEALTHIM_SERVER_URL")
	if serverURL == "" {
		serverURL = "https://stim.cxykevin.top" // Default server
	}
	
	server := NewServer(serverURL)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	// This should work with a real server
	err := server.Ping(ctx)
	if err != nil {
		t.Errorf("Ping failed: %v", err)
	}
}

// TestNewServer tests the NewServer function
func TestNewServer(t *testing.T) {
	server := NewServer("https://example.com")
	if server.client.BaseURL != "https://example.com" {
		t.Error("BaseURL not set correctly")
	}
	
	if server.client.HTTPClient == nil {
		t.Error("HTTPClient not initialized")
	}
}

// TestNewServerFromEnv tests the NewServerFromEnv function
func TestNewServerFromEnv(t *testing.T) {
	// Save original environment variable
	originalURL := os.Getenv("STEALTHIM_SERVER_URL")
	defer func() {
		if originalURL != "" {
			os.Setenv("STEALTHIM_SERVER_URL", originalURL)
		} else {
			os.Unsetenv("STEALTHIM_SERVER_URL")
		}
	}()
	
	// Set test environment variable
	testURL := "https://test.example.com"
	os.Setenv("STEALTHIM_SERVER_URL", testURL)
	
	server := NewServerFromEnv()
	if server.client.BaseURL != testURL {
		t.Error("BaseURL not set correctly from environment variable")
	}
}

// TestServerRegister tests the Server.Register method
func TestServerRegister(t *testing.T) {
	serverURL := os.Getenv("STEALTHIM_SERVER_URL")
	if serverURL == "" {
		serverURL = "https://stim.cxykevin.top" // Default server
	}
	
	server := NewServer(serverURL)
	username := generateRandomUsername()
	
	// Register a new user
	_, err := server.Register(context.Background(), username, "Ab123456", "Test User", username+"@example.com", "1234567890")
	if err != nil {
		t.Errorf("Failed to register user: %v", err)
	}
}

// TestServerLogin tests the Server.Login method
func TestServerLogin(t *testing.T) {
	serverURL := os.Getenv("STEALTHIM_SERVER_URL")
	if serverURL == "" {
		serverURL = "https://stim.cxykevin.top" // Default server
	}
	
	// First register a user
	username := generateRandomUsername()
	server := NewServer(serverURL)
	
	_, err := server.Register(context.Background(), username, "Ab123456", "Test User", username+"@example.com", "1234567890")
	if err != nil {
		t.Fatalf("Failed to register user: %v", err)
	}
	
	_, _, err = server.Login(context.Background(), username, "Ab123456")
	if err != nil {
		t.Fatalf("Failed to login: %v", err)
	}
}

// TestServerPingContextCancellation tests the Server.Ping method with context cancellation
func TestServerPingContextCancellation(t *testing.T) {
	serverURL := os.Getenv("STEALTHIM_SERVER_URL")
	if serverURL == "" {
		serverURL = "https://stim.cxykevin.top" // Default server
	}
	
	server := NewServer(serverURL)
	
	// Test context cancellation
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	
	// This should fail due to cancelled context
	err := server.Ping(ctx)
	if err == nil {
		t.Error("Expected error for cancelled context")
	}
}