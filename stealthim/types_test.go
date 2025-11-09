package stealthim

import (
	"testing"
)

// TestMessageTypeConstants tests the MessageType constants
func TestMessageTypeConstants(t *testing.T) {
	if int(Text) != 0 {
		t.Error("Text message type should be 0")
	}
	
	if int(Image) != 1 {
		t.Error("Image message type should be 1")
	}
	
	if int(Recall_Text) != 16 {
		t.Error("Recall_Text message type should be 16")
	}
}

// TestGroupMemberTypeConstants tests the GroupMemberType constants
func TestGroupMemberTypeConstants(t *testing.T) {
	if int(Member) != 0 {
		t.Error("Member type should be 0")
	}
	
	if int(Manager) != 1 {
		t.Error("Manager type should be 1")
	}
	
	if int(Owner) != 2 {
		t.Error("Owner type should be 2")
	}
}

// TestFileMetadataStructure tests the FileMetadata structure
func TestFileMetadataStructure(t *testing.T) {
	metadata := FileMetadata{
		Size:     1024,
		GroupID:  1,
		Hash:     "testhash",
		Filename: "test.txt",
	}
	
	if metadata.Size != 1024 {
		t.Error("Size not set correctly")
	}
	
	if metadata.GroupID != 1 {
		t.Error("GroupID not set correctly")
	}
	
	if metadata.Hash != "testhash" {
		t.Error("Hash not set correctly")
	}
	
	if metadata.Filename != "test.txt" {
		t.Error("Filename not set correctly")
	}
}