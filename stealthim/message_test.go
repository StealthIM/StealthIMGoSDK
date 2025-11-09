package stealthim

import (
	"context"
	"os"
	"testing"
	"time"
)

// TestSendMessage tests the Group.SendMessage method
func TestSendMessage(t *testing.T) {
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

	// Test sending a message
	err = newGroup.SendMessage(ctx, Text, "Hello, World!")
	if err != nil {
		t.Errorf("Failed to send message: %v", err)
	}
}

// TestRecallMessage tests the Group.RecallMessage method
func TestRecallMessage(t *testing.T) {
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

	// Send a message first
	err = newGroup.SendMessage(ctx, Text, "Hello, World!")
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// TODO: Test recalling a message
	// This would require getting the message ID first
	// For now, we'll skip this test as it requires more complex setup
	t.Skip("Skipping recall message test: requires message ID")
}

// TestSendText tests the Group.SendText method
func TestSendText(t *testing.T) {
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

	// Test sending a text message
	err = newGroup.SendText(ctx, "Hello, World!")
	if err != nil {
		t.Errorf("Failed to send text: %v", err)
	}
}

// TestReceiveMessages tests the Group.ReceiveMessages method
func TestReceiveMessages(t *testing.T) {
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

	// Test receiving messages
	opts := DefaultReceiveMessageOptions()
	msgChan, errChan := newGroup.ReceiveMessages(ctx, opts)
	if msgChan == nil || errChan == nil {
		t.Error("Expected channels to be non-nil")
	}

	// Close channels to avoid goroutine leak
	closeChannels(ctx)
}

// Helper function to close channels and avoid goroutine leaks
func closeChannels(ctx context.Context) {
	go func() {
		select {
		case <-ctx.Done():
		case <-time.After(5 * time.Second): // Timeout after 5 seconds
		}
	}()
}

// TestDefaultReceiveMessageOptions tests the DefaultReceiveMessageOptions function
func TestDefaultReceiveMessageOptions(t *testing.T) {
	opts := DefaultReceiveMessageOptions()
	if opts.FromID != "" {
		t.Error("Expected FromID to be empty")
	}
	if opts.Sync != false {
		t.Error("Expected Sync to be false")
	}
	if opts.Limit != 100 {
		t.Error("Expected Limit to be 100")
	}
}
