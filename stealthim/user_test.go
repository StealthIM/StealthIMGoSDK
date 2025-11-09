package stealthim

import (
	"context"
	"os"
	"testing"
	"time"
)

// TestUserGetSelfInfo tests the User.GetSelfInfo method
func TestUserGetSelfInfo(t *testing.T) {
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
	
	// Test getting self info
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	info, err := user.GetSelfInfo(ctx)
	if err != nil {
		t.Errorf("Failed to get self info: %v", err)
	}
	if info == nil {
		t.Error("Expected info to be non-nil")
	}
}

// TestUserGetUserInfo tests the User.GetUserInfo method
func TestUserGetUserInfo(t *testing.T) {
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
	
	// Test getting user info
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	info, err := user.GetUserInfo(ctx, username)
	if err != nil {
		t.Errorf("Failed to get user info: %v", err)
	}
	if info == nil {
		t.Error("Expected info to be non-nil")
	}
}

// TestUserChangePassword tests the User.ChangePassword method
func TestUserChangePassword(t *testing.T) {
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
	
	// Test changing password
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	err = user.ChangePassword(ctx, "Ab123456")
	if err != nil {
		t.Errorf("Failed to change password: %v", err)
	}
}

// TestUserChangeEmail tests the User.ChangeEmail method
func TestUserChangeEmail(t *testing.T) {
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
	
	// Test changing email
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	err = user.ChangeEmail(ctx, "newemail@example.com")
	if err != nil {
		t.Errorf("Failed to change email: %v", err)
	}
}

// TestUserChangeNickname tests the User.ChangeNickname method
func TestUserChangeNickname(t *testing.T) {
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
	
	// Test changing nickname
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	err = user.ChangeNickname(ctx, "New Nickname")
	if err != nil {
		t.Errorf("Failed to change nickname: %v", err)
	}
}

// TestUserChangePhoneNumber tests the User.ChangePhoneNumber method
func TestUserChangePhoneNumber(t *testing.T) {
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
	
	// Test changing phone number
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	err = user.ChangePhoneNumber(ctx, "0987654321")
	if err != nil {
		t.Errorf("Failed to change phone number: %v", err)
	}
}

// TestUserUpdateInfo tests the User.UpdateInfo method
func TestUserUpdateInfo(t *testing.T) {
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
	
	// Test updating info
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	err = user.UpdateInfo(ctx, "Ab123456", "newemail@example.com", "New Nickname", "0987654321")
	if err != nil {
		t.Errorf("Failed to update info: %v", err)
	}
}

// TestUserDelete tests the User.Delete method
func TestUserDelete(t *testing.T) {
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
	
	// Test deleting user
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	err = user.Delete(ctx)
	if err != nil {
		t.Errorf("Failed to delete user: %v", err)
	}
}

// TestUserGetGroups tests the User.GetGroups method
func TestUserGetGroups(t *testing.T) {
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
	
	// Test getting groups
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	groups, err := user.GetGroups(ctx)
	if err != nil {
		t.Errorf("Failed to get groups: %v", err)
	}
	if groups == nil {
		t.Error("Expected groups to be non-nil")
	}
}