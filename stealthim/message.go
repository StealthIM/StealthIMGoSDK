package stealthim

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// SendMessage sends a text message to the group
func (g *Group) SendMessage(ctx context.Context, msgType MessageType, content string) error {
	reqBody := map[string]any{
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
	reqBody := map[string]any{
		"type": int(RecallText),
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
	MsgID string // 消息ID，用于从特定消息开始拉取
}

// DefaultReceiveMessageOptions returns default options for receiving messages
func DefaultReceiveMessageOptions() *ReceiveMessageOptions {
	return &ReceiveMessageOptions{
		MsgID: "",
	}
}

// ReceiveMessages receives messages from the group via Server-Sent Events (SSE)
func (g *Group) ReceiveMessages(ctx context.Context, opts *ReceiveMessageOptions) (<-chan Message, <-chan error) {
	messageChan := make(chan Message)
	errorChan := make(chan error, 1)

	// Build query parameters
	queryParams := ""
	if opts.MsgID != "" {
		queryParams = fmt.Sprintf("?msgid=%s", opts.MsgID)
	}

	endpoint := fmt.Sprintf("/api/v1/message/%d%s", g.GroupID, queryParams)

	go func() {
		defer close(messageChan)
		defer close(errorChan)

		// Retry logic: attempt up to 3 times
		maxRetries := 3
		for attempt := 0; attempt < maxRetries; attempt++ {
			// Check if context was cancelled
			select {
			case <-ctx.Done():
				errorChan <- ctx.Err()
				return
			default:
			}

			// Create HTTP request
			req, err := http.NewRequestWithContext(ctx, "GET", g.client.BaseURL+endpoint, nil)
			if err != nil {
				errorChan <- fmt.Errorf("failed to create request: %w", err)
				return
			}

			// Set authorization header
			if g.client.Session != "" {
				req.Header.Set("Authorization", "Bearer "+g.client.Session)
			}

			// Set Accept header for SSE
			req.Header.Set("Accept", "text/event-stream")
			req.Header.Set("Cache-Control", "no-cache")
			req.Header.Set("Connection", "keep-alive")

			// Execute request
			resp, err := g.client.HTTPClient.Do(req)
			if err != nil {
				// If this was the last attempt, return the error
				if attempt == maxRetries-1 {
					errorChan <- fmt.Errorf("failed to execute request after %d attempts: %w", maxRetries, err)
					return
				}
				// Otherwise, continue to next attempt
				continue
			}

			// Check response status
			if resp.StatusCode != http.StatusOK {
				resp.Body.Close()
				// If this was the last attempt, return the error
				if attempt == maxRetries-1 {
					errorChan <- fmt.Errorf("request failed with status: %d after %d attempts", resp.StatusCode, maxRetries)
					return
				}
				// Otherwise, continue to next attempt
				continue
			}

			// Create a buffered reader for the response body
			reader := bufio.NewReader(resp.Body)

			// Process SSE stream
			for {
				// Check if context was cancelled
				select {
				case <-ctx.Done():
					resp.Body.Close()
					errorChan <- ctx.Err()
					return
				default:
				}

				// Read a line from the stream
				line, err := reader.ReadString('\n')
				if err != nil {
					resp.Body.Close()
					// Check if it's EOF (end of stream) - this might be normal
					if err.Error() == "EOF" {
						// For SSE streams, EOF can be normal after receiving all messages
						// If we've reached max attempts and still getting EOF, return normally
						if attempt == maxRetries-1 {
							return // Normal end of stream
						} else {
							// Retry the connection
							break // Break inner loop to retry connection
						}
					}
					// If this was the last attempt, return the error
					if attempt == maxRetries-1 {
						errorChan <- fmt.Errorf("failed to read from stream after %d attempts: %w", maxRetries, err)
						return
					}
					// Otherwise, continue to next attempt
					break // Break inner loop to retry connection
				}

				// Trim whitespace
				line = strings.TrimSpace(line)

				// Skip empty lines and comments
				if line == "" || strings.HasPrefix(line, ":") {
					continue
				}

				// Handle SSE data field
				if strings.HasPrefix(line, "data: ") {
					data := strings.TrimPrefix(line, "data: ")

					// Parse the JSON response
					var response struct {
						Result Result    `json:"result"`
						Msg    []Message `json:"msg"`
					}

					if err := json.Unmarshal([]byte(data), &response); err != nil {
						errorChan <- fmt.Errorf("failed to parse message: %w", err)
						continue
					}

					// Check if the response is successful
					if !response.Result.IsSuccess() {
						resp.Body.Close()
						errorChan <- response.Result.ToError()
						return
					}

					// Send each message to the channel
					for _, msg := range response.Msg {
						select {
						case messageChan <- msg:
						case <-ctx.Done():
							resp.Body.Close()
							errorChan <- ctx.Err()
							return
						}
					}
				}
			}

			// If we reach this point, we had a connection issue and need to retry
			// Wait briefly before retrying (except on the last attempt)
			if attempt < maxRetries-1 {
				select {
				case <-time.After(1 * time.Second): // Wait 1 second before retry
				case <-ctx.Done():
					errorChan <- ctx.Err()
					return
				}
			}
			// Continue to the next attempt in the loop
		}
	}()

	return messageChan, errorChan
}

// SendText sends a text message to the group
func (g *Group) SendText(ctx context.Context, message string) error {
	return g.SendMessage(ctx, Text, message)
}
