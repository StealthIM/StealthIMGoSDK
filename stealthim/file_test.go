package stealthim

import (
	"context"
	"os"
	"testing"
	"time"
)

// TestSendFile tests the Group.SendFile method
func TestSendFile(t *testing.T) {
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

	user, _, err := server.Login(context.Background(), username, "Ab123456")
	if err != nil {
		t.Fatalf("Failed to login: %v", err)
	}

	// Create a group
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	group := &Group{}
	newGroup, err := group.Create(ctx, user, "Test Group")
	if err != nil {
		t.Fatalf("Failed to create group: %v", err)
	}
	_ = newGroup

	// Create a temporary file for testing
	tmpFile, err := os.CreateTemp("", "test_file_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up after test

	// Write some test data to the file
	testData := "This is test data for file transfer"
	_, err = tmpFile.WriteString(testData)
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tmpFile.Close()

	// Test sending the file
	// Note: This test requires a working server connection
	// It may fail if the server is not accessible or has issues
	t.Logf("Attempting to send file to group %d", newGroup.GroupID)
	err = newGroup.SendFile(ctx, "test_file.txt", tmpFile.Name())
	if err != nil {
		// We're not asserting failure here because the test might fail due to server issues
		// rather than code issues
		t.Logf("SendFile returned error (may be due to server issues): %v", err)
		// 报告测试失败
		t.Fail()
	}
}

// TestCalculateBlake3Hash tests the calculateBlake3Hash function
func TestCalculateBlake3Hash(t *testing.T) {
	// Create a temporary file for testing
	tmpFile, err := os.CreateTemp("", "hash_test_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up after test

	// Write some test data to the file
	testData := "This is test data for hash calculation"
	_, err = tmpFile.WriteString(testData)
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tmpFile.Close()

	// Test the hash calculation function
	hash, err := calculateBlake3Hash(tmpFile.Name())
	if err != nil {
		t.Errorf("calculateBlake3Hash failed: %v", err)
	}

	// Verify that we got a hash result
	if hash == "" {
		t.Error("Expected non-empty hash")
	}

	// Verify hash format (should be hex string)
	if len(hash) != 64 {
		t.Errorf("Expected hash length of 64, got %d", len(hash))
	}
}

// TestClientDownloadFile tests the Client.DownloadFile method
func TestClientDownloadFile(t *testing.T) {
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

	user, _, err := server.Login(context.Background(), username, "Ab123456")
	if err != nil {
		t.Fatalf("Failed to login: %v", err)
	}

	// Test downloading a file (with a dummy hash)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// This will likely fail with a "file not found" error, which is expected
	// We're just testing that the method can be called
	err = user.client.DownloadFile(ctx, "dummyhash", "/tmp/output.txt", 4)
	if err == nil {
		t.Error("Expected error for dummy hash")
	}
}

// TestClientGetFileInfo tests the Client.GetFileInfo method
func TestClientGetFileInfo(t *testing.T) {
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

	user, _, err := server.Login(context.Background(), username, "Ab123456")
	if err != nil {
		t.Fatalf("Failed to login: %v", err)
	}

	// Test getting file info (with a dummy hash)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// This will likely fail with a "file not found" error, which is expected
	// We're just testing that the method can be called
	size, err := user.client.GetFileInfo(ctx, "dummyhash")
	if err == nil {
		t.Error("Expected error for dummy hash")
	}
	if size != 0 {
		t.Error("Expected size to be 0 for non-existent file")
	}
}
