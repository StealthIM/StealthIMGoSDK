package stealthim

import (
	"context"
	"os"
	"testing"
	"time"
)

// TestGroupCreate tests the Group.Create method
func TestGroupCreate(t *testing.T) {
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
	
	// Test creating a group
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	group := &Group{}
	newGroup, err := group.Create(ctx, user, "Test Group")
	if err != nil {
		t.Errorf("Failed to create group: %v", err)
	}
	if newGroup == nil {
		t.Error("Expected newGroup to be non-nil")
	}
}

// TestGroupJoin tests the Group.Join method
func TestGroupJoin(t *testing.T) {
	serverURL := os.Getenv("STEALTHIM_SERVER_URL")
	if serverURL == "" {
		serverURL = "https://stim.cxykevin.top" // Default server
	}
	
	// First register two users
	server := NewServer(serverURL)
	
	// Create first user (group owner)
	username1 := generateRandomUsername()
	_, err := server.Register(context.Background(), username1, "Ab123456", "Test User 1", username1+"@example.com", "1234567890")
	if err != nil {
		t.Fatalf("Failed to register user 1: %v", err)
	}
	
	user1, _, err := server.Login(context.Background(), username1, "Ab123456")
	if err != nil {
		t.Fatalf("Failed to login user 1: %v", err)
	}
	
	// Create second user (group member)
	username2 := generateRandomUsername()
	_, err = server.Register(context.Background(), username2, "Ab123456", "Test User 2", username2+"@example.com", "0987654321")
	if err != nil {
		t.Fatalf("Failed to register user 2: %v", err)
	}
	
	user2, _, err := server.Login(context.Background(), username2, "Ab123456")
	if err != nil {
		t.Fatalf("Failed to login user 2: %v", err)
	}
	_ = user2
	
	// Create a group with user1
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	group := &Group{}
	newGroup, err := group.Create(ctx, user1, "Test Group")
	if err != nil {
		t.Fatalf("Failed to create group: %v", err)
	}
	
	// Test joining the group with user2
	err = newGroup.Join(ctx, "")
	if err != nil {
		t.Errorf("Failed to join group: %v", err)
	}
}

// TestGroupGetMembers tests the Group.GetMembers method
func TestGroupGetMembers(t *testing.T) {
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
	
	// Test getting members
	members, err := newGroup.GetMembers(ctx)
	if err != nil {
		t.Errorf("Failed to get members: %v", err)
	}
	if members == nil {
		t.Error("Expected members to be non-nil")
	}
}

// TestGroupGetInfo tests the Group.GetInfo method
func TestGroupGetInfo(t *testing.T) {
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
	
	// Test getting group info
	info, err := newGroup.GetInfo(ctx)
	if err != nil {
		t.Errorf("Failed to get group info: %v", err)
	}
	if info == nil {
		t.Error("Expected info to be non-nil")
	}
}

// TestGroupInvite tests the Group.Invite method
func TestGroupInvite(t *testing.T) {
	serverURL := os.Getenv("STEALTHIM_SERVER_URL")
	if serverURL == "" {
		serverURL = "https://stim.cxykevin.top" // Default server
	}
	
	// First register two users
	server := NewServer(serverURL)
	
	// Create first user (group owner)
	username1 := generateRandomUsername()
	_, err := server.Register(context.Background(), username1, "Ab123456", "Test User 1", username1+"@example.com", "1234567890")
	if err != nil {
		t.Fatalf("Failed to register user 1: %v", err)
	}
	
	user1, _, err := server.Login(context.Background(), username1, "Ab123456")
	if err != nil {
		t.Fatalf("Failed to login user 1: %v", err)
	}
	
	// Create second user (invitee)
	username2 := generateRandomUsername()
	_, err = server.Register(context.Background(), username2, "Ab123456", "Test User 2", username2+"@example.com", "0987654321")
	if err != nil {
		t.Fatalf("Failed to register user 2: %v", err)
	}
	
	// Create a group with user1
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	group := &Group{}
	newGroup, err := group.Create(ctx, user1, "Test Group")
	if err != nil {
		t.Fatalf("Failed to create group: %v", err)
	}
	
	// Test inviting user2 to the group
	err = newGroup.Invite(ctx, username2)
	if err != nil {
		t.Errorf("Failed to invite user: %v", err)
	}
}

// TestGroupSetMemberRole tests the Group.SetMemberRole method
func TestGroupSetMemberRole(t *testing.T) {
	serverURL := os.Getenv("STEALTHIM_SERVER_URL")
	if serverURL == "" {
		serverURL = "https://stim.cxykevin.top" // Default server
	}
	
	// First register two users
	server := NewServer(serverURL)
	
	// Create first user (group owner)
	username1 := generateRandomUsername()
	_, err := server.Register(context.Background(), username1, "Ab123456", "Test User 1", username1+"@example.com", "1234567890")
	if err != nil {
		t.Fatalf("Failed to register user 1: %v", err)
	}
	
	user1, _, err := server.Login(context.Background(), username1, "Ab123456")
	if err != nil {
		t.Fatalf("Failed to login user 1: %v", err)
	}
	
	// Create second user (group member)
	username2 := generateRandomUsername()
	_, err = server.Register(context.Background(), username2, "Ab123456", "Test User 2", username2+"@example.com", "0987654321")
	if err != nil {
		t.Fatalf("Failed to register user 2: %v", err)
	}
	
	user2, _, err := server.Login(context.Background(), username2, "Ab123456")
	if err != nil {
		t.Fatalf("Failed to login user 2: %v", err)
	}
	_ = user2
	
	// Create a group with user1
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	group := &Group{}
	newGroup, err := group.Create(ctx, user1, "Test Group")
	if err != nil {
		t.Fatalf("Failed to create group: %v", err)
	}
	
	// Add user2 to the group
	err = newGroup.Join(ctx, "")
	if err != nil {
		t.Fatalf("Failed to join group: %v", err)
	}
	
	// Test setting member role
	err = newGroup.SetMemberRole(ctx, username2, Manager)
	if err != nil {
		t.Errorf("Failed to set member role: %v", err)
	}
}

// TestGroupKick tests the Group.Kick method
func TestGroupKick(t *testing.T) {
	serverURL := os.Getenv("STEALTHIM_SERVER_URL")
	if serverURL == "" {
		serverURL = "https://stim.cxykevin.top" // Default server
	}
	
	// First register two users
	server := NewServer(serverURL)
	
	// Create first user (group owner)
	username1 := generateRandomUsername()
	_, err := server.Register(context.Background(), username1, "Ab123456", "Test User 1", username1+"@example.com", "1234567890")
	if err != nil {
		t.Fatalf("Failed to register user 1: %v", err)
	}
	
	user1, _, err := server.Login(context.Background(), username1, "Ab123456")
	if err != nil {
		t.Fatalf("Failed to login user 1: %v", err)
	}
	
	// Create second user (group member)
	username2 := generateRandomUsername()
	_, err = server.Register(context.Background(), username2, "Ab123456", "Test User 2", username2+"@example.com", "0987654321")
	if err != nil {
		t.Fatalf("Failed to register user 2: %v", err)
	}
	
	user2, _, err := server.Login(context.Background(), username2, "Ab123456")
	if err != nil {
		t.Fatalf("Failed to login user 2: %v", err)
	}
	_ = user2
	
	// Create a group with user1
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	group := &Group{}
	newGroup, err := group.Create(ctx, user1, "Test Group")
	if err != nil {
		t.Fatalf("Failed to create group: %v", err)
	}
	
	// Add user2 to the group
	err = newGroup.Join(ctx, "")
	if err != nil {
		t.Fatalf("Failed to join group: %v", err)
	}
	
	// Test kicking user2 from the group
	err = newGroup.Kick(ctx, username2)
	if err != nil {
		t.Errorf("Failed to kick user: %v", err)
	}
}

// TestGroupChangeName tests the Group.ChangeName method
func TestGroupChangeName(t *testing.T) {
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
	
	// Test changing group name
	err = newGroup.ChangeName(ctx, "New Group Name")
	if err != nil {
		t.Errorf("Failed to change group name: %v", err)
	}
}

// TestGroupChangePassword tests the Group.ChangePassword method
func TestGroupChangePassword(t *testing.T) {
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
	
	// Test changing group password
	err = newGroup.ChangePassword(ctx, "newpassword123")
	if err != nil {
		t.Errorf("Failed to change group password: %v", err)
	}
}