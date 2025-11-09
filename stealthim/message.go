package stealthim

import (
	"context"
	"fmt"
	"net/http"
)

// SendMessage sends a text message to the group
func (g *Group) SendMessage(ctx context.Context, msgType MessageType, content string) error {
	reqBody := map[string]interface{}{
		"type": int(msgType),
		"msg":  content,
	}

	endpoint := fmt.Sprintf("/api/v1/message/%d", g.GroupID)
	resp, err := g.client.doRequest(ctx, "POST", endpoint, reqBody)
	if err != nil {
		return fmt.Errorf("send message request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result Result `json:"result"`
	}
	if err := g.client.parseResponse(resp, &response); err != nil {
		return fmt.Errorf("failed to parse send message response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return response.Result.ToError()
	}

	return nil
}

// RecallMessage recalls/deletes a message
func (g *Group) RecallMessage(ctx context.Context, messageID string) error {
	reqBody := map[string]interface{}{
		"type": int(Recall_Text),
		"msg":  messageID,
	}

	endpoint := fmt.Sprintf("/api/v1/message/%d", g.GroupID)
	resp, err := g.client.doRequest(ctx, "POST", endpoint, reqBody)
	if err != nil {
		return fmt.Errorf("recall message request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result Result `json:"result"`
	}
	if err := g.client.parseResponse(resp, &response); err != nil {
		return fmt.Errorf("failed to parse recall message response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return response.Result.ToError()
	}

	return nil
}

// ReceiveMessageOptions holds options for receiving messages
type ReceiveMessageOptions struct {
	FromID string
	Sync   bool
	Limit  int
}

// DefaultReceiveMessageOptions returns default options for receiving messages
func DefaultReceiveMessageOptions() *ReceiveMessageOptions {
	return &ReceiveMessageOptions{
		FromID: "",
		Sync:   false,
		Limit:  100,
	}
}

// ReceiveMessages receives messages from the group via Server-Sent Events (SSE)
// This is a simplified implementation - a full implementation would require proper SSE handling
func (g *Group) ReceiveMessages(ctx context.Context, opts *ReceiveMessageOptions) (<-chan Message, <-chan error) {
	messageChan := make(chan Message)
	errorChan := make(chan error, 1)

	// Build query parameters
	queryParams := "?"
	if opts.FromID != "" {
		queryParams += fmt.Sprintf("from_id=%s&", opts.FromID)
	}
	queryParams += fmt.Sprintf("sync=%t&", opts.Sync)
	queryParams += fmt.Sprintf("limit=%d", opts.Limit)

	endpoint := fmt.Sprintf("/api/v1/message/%d%s", g.GroupID, queryParams)

	// In a real implementation, this would properly handle Server-Sent Events
	// For now, we'll provide a basic structure that shows the intended API
	go func() {
		defer close(messageChan)
		defer close(errorChan)

		// This is a simplified placeholder implementation
		// A full implementation would need to properly handle SSE streams
		req, err := http.NewRequestWithContext(ctx, "GET", g.client.BaseURL+endpoint, nil)
		if err != nil {
			errorChan <- fmt.Errorf("failed to create request: %w", err)
			return
		}

		// Set authorization header
		if g.client.Session != "" {
			req.Header.Set("Authorization", "Bearer "+g.client.Session)
		}

		// Note: This is simplified - a real implementation would need to properly handle SSE
		// For now, we'll just make a regular GET request to demonstrate the concept
		resp, err := g.client.HTTPClient.Do(req)
		if err != nil {
			errorChan <- fmt.Errorf("failed to execute request: %w", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			errorChan <- fmt.Errorf("request failed with status: %d", resp.StatusCode)
			return
		}

		// In a real implementation, we would parse the SSE stream here
		// This is just a placeholder to indicate the intended functionality
	}()

	return messageChan, errorChan
}

// SendText sends a text message to the group
func (g *Group) SendText(ctx context.Context, message string) error {
	return g.SendMessage(ctx, Text, message)
}