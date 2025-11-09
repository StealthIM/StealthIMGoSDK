package stealthim

import (
	"context"
	"fmt"
)

// User represents an authenticated user
type User struct {
	client *Client
	info   UserInfo
}

// GetSelfInfo retrieves the current user's information
func (u *User) GetSelfInfo(ctx context.Context) (*UserInfo, error) {
	resp, err := u.client.doRequest(ctx, "GET", "/api/v1/user", nil)
	if err != nil {
		return nil, fmt.Errorf("get self info request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result   Result   `json:"result"`
		UserInfo UserInfo `json:"user_info"`
	}
	if err := u.client.parseResponse(resp, &response); err != nil {
		return nil, fmt.Errorf("failed to parse get self info response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return nil, response.Result.ToError()
	}

	return &response.UserInfo, nil
}

// GetUserInfo retrieves another user's public information
func (u *User) GetUserInfo(ctx context.Context, username string) (*UserInfo, error) {
	endpoint := fmt.Sprintf("/api/v1/user/%s", username)
	resp, err := u.client.doRequest(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("get user info request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result   Result   `json:"result"`
		UserInfo UserInfo `json:"user_info"`
	}
	if err := u.client.parseResponse(resp, &response); err != nil {
		return nil, fmt.Errorf("failed to parse get user info response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return nil, response.Result.ToError()
	}

	return &response.UserInfo, nil
}

// ChangePassword updates the user's password
func (u *User) ChangePassword(ctx context.Context, newPassword string) error {
	reqBody := map[string]interface{}{
		"password": newPassword,
	}

	resp, err := u.client.doRequest(ctx, "PUT", "/api/v1/user/password", reqBody)
	if err != nil {
		return fmt.Errorf("change password request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result Result `json:"result"`
	}
	if err := u.client.parseResponse(resp, &response); err != nil {
		return fmt.Errorf("failed to parse change password response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return response.Result.ToError()
	}

	return nil
}

// ChangeEmail updates the user's email
func (u *User) ChangeEmail(ctx context.Context, newEmail string) error {
	reqBody := map[string]interface{}{
		"email": newEmail,
	}

	resp, err := u.client.doRequest(ctx, "PUT", "/api/v1/user/email", reqBody)
	if err != nil {
		return fmt.Errorf("change email request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result Result `json:"result"`
	}
	if err := u.client.parseResponse(resp, &response); err != nil {
		return fmt.Errorf("failed to parse change email response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return response.Result.ToError()
	}

	return nil
}

// ChangeNickname updates the user's nickname
func (u *User) ChangeNickname(ctx context.Context, newNickname string) error {
	reqBody := map[string]interface{}{
		"nickname": newNickname,
	}

	resp, err := u.client.doRequest(ctx, "PUT", "/api/v1/user/nickname", reqBody)
	if err != nil {
		return fmt.Errorf("change nickname request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result Result `json:"result"`
	}
	if err := u.client.parseResponse(resp, &response); err != nil {
		return fmt.Errorf("failed to parse change nickname response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return response.Result.ToError()
	}

	return nil
}

// ChangePhoneNumber updates the user's phone number
func (u *User) ChangePhoneNumber(ctx context.Context, newPhoneNumber string) error {
	reqBody := map[string]interface{}{
		"phone_number": newPhoneNumber,
	}

	resp, err := u.client.doRequest(ctx, "PUT", "/api/v1/user/phone", reqBody)
	if err != nil {
		return fmt.Errorf("change phone number request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result Result `json:"result"`
	}
	if err := u.client.parseResponse(resp, &response); err != nil {
		return fmt.Errorf("failed to parse change phone number response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return response.Result.ToError()
	}

	return nil
}

// UpdateInfo updates multiple user fields at once
func (u *User) UpdateInfo(ctx context.Context, password, email, nickname, phoneNumber string) error {
	reqBody := map[string]interface{}{}
	if password != "" {
		reqBody["password"] = password
	}
	if email != "" {
		reqBody["email"] = email
	}
	if nickname != "" {
		reqBody["nickname"] = nickname
	}
	if phoneNumber != "" {
		reqBody["phone_number"] = phoneNumber
	}

	resp, err := u.client.doRequest(ctx, "PUT", "/api/v1/user", reqBody)
	if err != nil {
		return fmt.Errorf("update info request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result Result `json:"result"`
	}
	if err := u.client.parseResponse(resp, &response); err != nil {
		return fmt.Errorf("failed to parse update info response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return response.Result.ToError()
	}

	return nil
}

// Delete deletes the user account
func (u *User) Delete(ctx context.Context) error {
	resp, err := u.client.doRequest(ctx, "DELETE", "/api/v1/user", nil)
	if err != nil {
		return fmt.Errorf("delete user request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result Result `json:"result"`
	}
	if err := u.client.parseResponse(resp, &response); err != nil {
		return fmt.Errorf("failed to parse delete user response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return response.Result.ToError()
	}

	return nil
}

// GetGroups retrieves the user's groups
func (u *User) GetGroups(ctx context.Context) ([]int64, error) {
	resp, err := u.client.doRequest(ctx, "GET", "/api/v1/group", nil)
	if err != nil {
		return nil, fmt.Errorf("get groups request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result Result  `json:"result"`
		Groups []int64 `json:"groups"`
	}
	if err := u.client.parseResponse(resp, &response); err != nil {
		return nil, fmt.Errorf("failed to parse get groups response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return nil, response.Result.ToError()
	}

	return response.Groups, nil
}