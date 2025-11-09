package stealthim

import (
	"context"
	"fmt"
)

// Group represents a group in StealthIM
type Group struct {
	client  *Client
	GroupID int64
}

// Create creates a new group
func (g *Group) Create(ctx context.Context, user *User, groupName string) (*Group, error) {
	reqBody := map[string]interface{}{
		"name": groupName,
	}

	resp, err := user.client.doRequest(ctx, "POST", "/api/v1/group", reqBody)
	if err != nil {
		return nil, fmt.Errorf("create group request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result  Result `json:"result"`
		GroupID int64  `json:"groupid"`
	}
	if err := user.client.parseResponse(resp, &response); err != nil {
		return nil, fmt.Errorf("failed to parse create group response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return nil, response.Result.ToError()
	}

	return &Group{
		client:  user.client,
		GroupID: response.GroupID,
	}, nil
}

// Join joins a group with a password
func (g *Group) Join(ctx context.Context, password string) error {
	reqBody := map[string]interface{}{
		"password": password,
	}

	endpoint := fmt.Sprintf("/api/v1/group/%d", g.GroupID)
	resp, err := g.client.doRequest(ctx, "POST", endpoint, reqBody)
	if err != nil {
		return fmt.Errorf("join group request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result Result `json:"result"`
	}
	if err := g.client.parseResponse(resp, &response); err != nil {
		return fmt.Errorf("failed to parse join group response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return response.Result.ToError()
	}

	return nil
}

// GetMembers retrieves the group member list
func (g *Group) GetMembers(ctx context.Context) ([]GroupMember, error) {
	endpoint := fmt.Sprintf("/api/v1/group/%d", g.GroupID)
	resp, err := g.client.doRequest(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("get members request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result  Result        `json:"result"`
		Members []GroupMember `json:"members"`
	}
	if err := g.client.parseResponse(resp, &response); err != nil {
		return nil, fmt.Errorf("failed to parse get members response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return nil, response.Result.ToError()
	}

	return response.Members, nil
}

// GetInfo retrieves public group information
func (g *Group) GetInfo(ctx context.Context) ([]GroupMember, error) {
	endpoint := fmt.Sprintf("/api/v1/group/%d/public", g.GroupID)
	resp, err := g.client.doRequest(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("get group info request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result  Result        `json:"result"`
		Members []GroupMember `json:"members"`
	}
	if err := g.client.parseResponse(resp, &response); err != nil {
		return nil, fmt.Errorf("failed to parse get group info response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return nil, response.Result.ToError()
	}

	return response.Members, nil
}

// Invite invites a user to the group
func (g *Group) Invite(ctx context.Context, username string) error {
	reqBody := map[string]interface{}{
		"username": username,
	}

	endpoint := fmt.Sprintf("/api/v1/group/%d/invite", g.GroupID)
	resp, err := g.client.doRequest(ctx, "POST", endpoint, reqBody)
	if err != nil {
		return fmt.Errorf("invite user request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result Result `json:"result"`
	}
	if err := g.client.parseResponse(resp, &response); err != nil {
		return fmt.Errorf("failed to parse invite user response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return response.Result.ToError()
	}

	return nil
}

// SetMemberRole sets a user's role in the group
func (g *Group) SetMemberRole(ctx context.Context, username string, role GroupMemberType) error {
	reqBody := map[string]interface{}{
		"username": username,
		"type":     int(role),
	}

	endpoint := fmt.Sprintf("/api/v1/group/%d/role", g.GroupID)
	resp, err := g.client.doRequest(ctx, "PUT", endpoint, reqBody)
	if err != nil {
		return fmt.Errorf("set member role request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result Result `json:"result"`
	}
	if err := g.client.parseResponse(resp, &response); err != nil {
		return fmt.Errorf("failed to parse set member role response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return response.Result.ToError()
	}

	return nil
}

// Kick removes a user from the group
func (g *Group) Kick(ctx context.Context, username string) error {
	reqBody := map[string]interface{}{
		"username": username,
	}

	endpoint := fmt.Sprintf("/api/v1/group/%d/kick", g.GroupID)
	resp, err := g.client.doRequest(ctx, "POST", endpoint, reqBody)
	if err != nil {
		return fmt.Errorf("kick user request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result Result `json:"result"`
	}
	if err := g.client.parseResponse(resp, &response); err != nil {
		return fmt.Errorf("failed to parse kick user response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return response.Result.ToError()
	}

	return nil
}

// ChangeName changes the group name
func (g *Group) ChangeName(ctx context.Context, newName string) error {
	reqBody := map[string]interface{}{
		"name": newName,
	}

	endpoint := fmt.Sprintf("/api/v1/group/%d/name", g.GroupID)
	resp, err := g.client.doRequest(ctx, "PUT", endpoint, reqBody)
	if err != nil {
		return fmt.Errorf("change group name request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result Result `json:"result"`
	}
	if err := g.client.parseResponse(resp, &response); err != nil {
		return fmt.Errorf("failed to parse change group name response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return response.Result.ToError()
	}

	return nil
}

// ChangePassword changes the group password
func (g *Group) ChangePassword(ctx context.Context, newPassword string) error {
	reqBody := map[string]interface{}{
		"password": newPassword,
	}

	endpoint := fmt.Sprintf("/api/v1/group/%d/password", g.GroupID)
	resp, err := g.client.doRequest(ctx, "PUT", endpoint, reqBody)
	if err != nil {
		return fmt.Errorf("change group password request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result Result `json:"result"`
	}
	if err := g.client.parseResponse(resp, &response); err != nil {
		return fmt.Errorf("failed to parse change group password response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return response.Result.ToError()
	}

	return nil
}