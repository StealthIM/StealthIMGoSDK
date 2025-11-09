package stealthim

import (
	"os"
	"testing"
)

// TestGenerateRandomString tests the GenerateRandomString function
func TestGenerateRandomString(t *testing.T) {
	str, err := GenerateRandomString(16)
	if err != nil {
		t.Errorf("Failed to generate random string: %v", err)
	}
	
	if len(str) != 16 {
		t.Errorf("Expected string length of 16, got %d", len(str))
	}
	
	// Generate another string and ensure it's different
	str2, err := GenerateRandomString(16)
	if err != nil {
		t.Errorf("Failed to generate random string: %v", err)
	}
	
	if str == str2 {
		t.Error("Expected two random strings to be different")
	}
}

// TestReadFile tests the ReadFile function
func TestReadFile(t *testing.T) {
	// Create a temporary file for testing
	content := []byte("Hello, World!")
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	
	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}
	
	// Test reading the file
	data, err := ReadFile(tmpfile.Name())
	if err != nil {
		t.Errorf("Failed to read file: %v", err)
	}
	
	if string(data) != string(content) {
		t.Errorf("Expected %s, got %s", string(content), string(data))
	}
}

// TestWriteFile tests the WriteFile function
func TestWriteFile(t *testing.T) {
	// Create a temporary file path for testing
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	
	// Close and remove the file to test writing to a new file
	tmpfile.Close()
	os.Remove(tmpfile.Name())
	
	content := []byte("Hello, World!")
	err = WriteFile(tmpfile.Name(), content)
	if err != nil {
		t.Errorf("Failed to write file: %v", err)
	}
	
	// Verify the file was written correctly
	data, err := ReadFile(tmpfile.Name())
	if err != nil {
		t.Errorf("Failed to read file: %v", err)
	}
	
	if string(data) != string(content) {
		t.Errorf("Expected %s, got %s", string(content), string(data))
	}
}

// TestFileExists tests the FileExists function
func TestFileExists(t *testing.T) {
	// Create a temporary file for testing
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	
	// Test that the file exists
	if !FileExists(tmpfile.Name()) {
		t.Error("Expected file to exist")
	}
	
	// Test that a non-existent file doesn't exist
	if FileExists("/non/existent/file") {
		t.Error("Expected file to not exist")
	}
}