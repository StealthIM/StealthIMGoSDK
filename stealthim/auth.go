package stealthim

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Server represents the StealthIM server
type Server struct {
	client *Client
}

// NewServer creates a new server instance
func NewServer(baseURL string) *Server {
	return &Server{
		client: NewClient(baseURL),
	}
}

// NewServerFromEnv creates a new server instance using the STEALTHIM_SERVER_URL environment variable
func NewServerFromEnv() *Server {
	serverURL := os.Getenv("STEALTHIM_SERVER_URL")
	return &Server{
		client: NewClient(serverURL),
	}
}

// Ping checks if the server is available
func (s *Server) Ping(ctx context.Context) error {
	resp, err := s.client.doRequest(ctx, "GET", "/api/v1/ping", nil)
	if err != nil {
		return fmt.Errorf("ping request failed: %w", err)
	}
	defer resp.Body.Close()

	// Ping API returns a simple message, not a Result structure
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ping request failed with status: %d", resp.StatusCode)
	}

	// Read the response body to check if it's valid
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read ping response: %w", err)
	}

	// Check if the response contains the expected message
	var pingResp map[string]any
	if err := json.Unmarshal(body, &pingResp); err != nil {
		return fmt.Errorf("failed to parse ping response: %w", err)
	}

	if msg, ok := pingResp["message"]; !ok {
		return fmt.Errorf("ping response does not contain message field")
	} else if msgStr, ok := msg.(string); !ok || msgStr == "" {
		return fmt.Errorf("ping message is not a valid string")
	}

	return nil
}

// Register registers a new user
func (s *Server) Register(ctx context.Context, username, password, nickname, email, phoneNumber string) (*UserInfo, error) {
	reqBody := map[string]any{
		"username":     username,
		"password":     password,
		"nickname":     nickname,
		"email":        email,
		"phone_number": phoneNumber,
	}

	resp, err := s.client.doRequest(ctx, "POST", "/api/v1/user/register", reqBody)
	if err != nil {
		return nil, fmt.Errorf("register request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result Result `json:"result"`
	}
	if err := s.client.parseResponse(resp, &response); err != nil {
		return nil, fmt.Errorf("failed to parse register response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return nil, response.Result.ToError()
	}

	// For register, we just return success since it doesn't return session/user_info like login
	return nil, nil
}

// Login authenticates a user and returns user information
func (s *Server) Login(ctx context.Context, username, password string) (*User, *UserInfo, error) {
	reqBody := map[string]any{
		"username": username,
		"password": password,
	}

	resp, err := s.client.doRequest(ctx, "POST", "/api/v1/user", reqBody)
	if err != nil {
		return nil, nil, fmt.Errorf("login request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result   Result   `json:"result"`
		Session  string   `json:"session"`
		UserInfo UserInfo `json:"user_info"`
	}
	if err := s.client.parseResponse(resp, &response); err != nil {
		return nil, nil, fmt.Errorf("failed to parse login response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return nil, nil, response.Result.ToError()
	}

	// Update client with session
	s.client.Session = response.Session

	user := &User{
		client: s.client,
		info:   response.UserInfo,
	}

	return user, &response.UserInfo, nil
}
