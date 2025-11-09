package stealthim

import (
	"context"
	"os"
	"testing"
	"time"
)

// TestResultIsSuccess tests the Result.IsSuccess method
func TestResultIsSuccess(t *testing.T) {
	result := Result{Code: 800, Msg: "success"}
	if !result.IsSuccess() {
		t.Error("Expected result to be successful")
	}

	result.Code = 900
	if result.IsSuccess() {
		t.Error("Expected result to be unsuccessful")
	}
}

// TestResultToError tests the Result.ToError method
func TestResultToError(t *testing.T) {
	result := Result{Code: 800, Msg: "success"}
	if err := result.ToError(); err != nil {
		t.Error("Expected no error for successful result")
	}

	result.Code = 900
	result.Msg = "server error"
	if err := result.ToError(); err == nil {
		t.Error("Expected error for unsuccessful result")
	}
}

// TestNewClient tests the NewClient function
func TestNewClient(t *testing.T) {
	client := NewClient("https://example.com")
	if client.BaseURL != "https://example.com" {
		t.Error("BaseURL not set correctly")
	}
	
	if client.HTTPClient == nil {
		t.Error("HTTPClient not initialized")
	}
}

// TestNewClientWithSession tests the NewClientWithSession function
func TestNewClientWithSession(t *testing.T) {
	client := NewClientWithSession("https://example.com", "test-session")
	if client.BaseURL != "https://example.com" {
		t.Error("BaseURL not set correctly")
	}
	
	if client.Session != "test-session" {
		t.Error("Session not set correctly")
	}
}

// TestNewClientFromEnv tests the NewClientFromEnv function
func TestNewClientFromEnv(t *testing.T) {
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
	
	client := NewClientFromEnv()
	if client.BaseURL != testURL {
		t.Error("BaseURL not set correctly from environment variable")
	}
}

// TestClientDoRequest tests the Client.doRequest method
func TestClientDoRequest(t *testing.T) {
	// Get server URL from environment variable
	serverURL := os.Getenv("STEALTHIM_SERVER_URL")
	
	client := NewClient(serverURL)
	
	// Test context cancellation
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	
	// This should fail due to timeout
	_, err := client.doRequest(ctx, "GET", "/api/v1/ping", nil)
	if err == nil {
		t.Error("Expected error for timeout")
	}
}