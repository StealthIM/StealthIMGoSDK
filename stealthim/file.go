package stealthim

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
	"lukechampine.com/blake3"
)

// SendFile uploads a file to the group using WebSocket
func (g *Group) SendFile(ctx context.Context, filename, filepath string) error {
	// Connect to the WebSocket endpoint
	// Convert HTTP/HTTPS URL to WebSocket URL
	var wsURL string
	if g.client.BaseURL[:5] == "https" {
		wsURL = "wss" + g.client.BaseURL[5:] + "/api/v1/file"
	} else {
		wsURL = "ws" + g.client.BaseURL[4:] + "/api/v1/file"
	}
	
	// Add authorization as query parameter if available
	if g.client.Session != "" {
		// Parse URL and add authorization parameter
		u, err := url.Parse(wsURL)
		if err != nil {
			return fmt.Errorf("failed to parse WebSocket URL: %w", err)
		}
		
		// Add authorization parameter
		q := u.Query()
		q.Set("authorization", g.client.Session)
		u.RawQuery = q.Encode()
		
		wsURL = u.String()
	}

	// Attempt to connect to the WebSocket
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		return fmt.Errorf("failed to connect to WebSocket: %w", err)
	}
	defer conn.Close()

	// Open the file
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Get file info
	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}
	fileSize := fileInfo.Size()

	// Calculate hash using Blake3 algorithm
	// The algorithm should split the file into 2048 KiB blocks and hash each block
	// then concatenate the binary hash results and hash again
	hash, err := calculateBlake3Hash(filepath)
	if err != nil {
		return fmt.Errorf("failed to calculate file hash: %w", err)
	}

	// Prepare metadata
	metadata := FileMetadata{
		Size:     fileSize,
		GroupID:  g.GroupID,
		Hash:     hash,
		Filename: filename,
	}

	// Send metadata
	if err := conn.WriteJSON(metadata); err != nil {
		return fmt.Errorf("failed to send metadata: %w", err)
	}

	// Read response for metadata
	var metaResponse struct {
		Result struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		} `json:"result"`
		Type string `json:"type"`
	}
	if err := conn.ReadJSON(&metaResponse); err != nil {
		return fmt.Errorf("failed to read metadata response: %w", err)
	}

	if metaResponse.Result.Code != 800 {
		return fmt.Errorf("metadata upload failed: %s", metaResponse.Result.Msg)
	}

	// Upload file in chunks
	blockSize := int64(2048 * 1024) // 2048 KiB
	buffer := make([]byte, blockSize)
	blockID := int32(0)

	for {
		// Read a block from the file
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return fmt.Errorf("failed to read file: %w", err)
		}
		if n == 0 {
			break // End of file
		}

		// Prepare the block data: 4-byte block ID + block content
		blockData := make([]byte, 4+n)
		binary.LittleEndian.PutUint32(blockData[:4], uint32(blockID))
		copy(blockData[4:], buffer[:n])

		// Send the block via WebSocket binary message
		if err := conn.WriteMessage(websocket.BinaryMessage, blockData); err != nil {
			return fmt.Errorf("failed to send block: %w", err)
		}

		// Read response for this block
		var blockResponse struct {
			Result struct {
				Code int    `json:"code"`
				Msg  string `json:"msg"`
			} `json:"result"`
			Type    string `json:"type"`
			BlockID int32  `json:"blockid"`
		}
		if err := conn.ReadJSON(&blockResponse); err != nil {
			return fmt.Errorf("failed to read block response: %w", err)
		}

		if blockResponse.Result.Code != 800 {
			return fmt.Errorf("block upload failed: %s", blockResponse.Result.Msg)
		}

		blockID++
	}

	// Wait for completion response
	var completeResponse struct {
		Result struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		} `json:"result"`
		Type string `json:"type"`
	}
	if err := conn.ReadJSON(&completeResponse); err != nil {
		return fmt.Errorf("failed to read completion response: %w", err)
	}

	if completeResponse.Result.Code != 800 {
		return fmt.Errorf("file upload failed: %s", completeResponse.Result.Msg)
	}

	return nil
}

// calculateBlake3Hash calculates the Blake3 hash according to the specified algorithm
func calculateBlake3Hash(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Get file info to determine size
	_, err = file.Stat()
	if err != nil {
		return "", fmt.Errorf("failed to get file info: %w", err)
	}

	blockSize := int64(2048 * 1024) // 2048 KiB
	var hashes []byte

	// Process the file in blocks
	buffer := make([]byte, blockSize)
	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return "", fmt.Errorf("failed to read file: %w", err)
		}
		if n == 0 {
			break
		}

		// Calculate hash for this block
		blockHash := blake3.Sum256(buffer[:n])
		hashes = append(hashes, blockHash[:]...)
	}

	// Calculate final hash from concatenated block hashes
	finalHash := blake3.Sum256(hashes)
	return fmt.Sprintf("%x", finalHash[:]), nil
}

// DownloadFile downloads a file with multi-threading support
// This is a placeholder implementation - a full implementation would handle the range requests properly
func (c *Client) DownloadFile(ctx context.Context, fileHash, outputPath string, threads int) error {
	// For now, implement a basic download
	// The API supports Range header for partial downloads
	endpoint := fmt.Sprintf("/api/v1/file/%s", fileHash)

	resp, err := c.doRequest(ctx, "GET", endpoint, nil)
	if err != nil {
		return fmt.Errorf("download request failed: %w", err)
	}
	defer resp.Body.Close()

	// Create output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	// Copy response body to file
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// GetFileInfo gets information about a file
func (c *Client) GetFileInfo(ctx context.Context, fileHash string) (int64, error) {
	endpoint := fmt.Sprintf("/api/v1/file/%s", fileHash)
	resp, err := c.doRequest(ctx, "GET", endpoint, nil)
	if err != nil {
		return 0, fmt.Errorf("get file info request failed: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Result Result `json:"result"`
		Size   int64  `json:"size"`
	}
	if err := c.parseResponse(resp, &response); err != nil {
		return 0, fmt.Errorf("failed to parse get file info response: %w", err)
	}

	if !response.Result.IsSuccess() {
		return 0, response.Result.ToError()
	}

	return response.Size, nil
}