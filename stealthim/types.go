package stealthim

// Result represents the API response result
type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// UserInfo represents user information
type UserInfo struct {
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	VIP         int    `json:"vip"`
	CreateTime  string `json:"create_time"`
}

// GroupMember represents a group member
type GroupMember struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}

// Message represents a message
type Message struct {
	GroupID  string `json:"groupid"`
	Msg      string `json:"msg"`
	MsgID    string `json:"msgid"`
	Time     string `json:"time"`
	Type     int    `json:"type"`
	Username string `json:"username"`
	Hash     string `json:"hash,omitempty"`
}

// MessageType represents the type of message
type MessageType int

// Message types
const (
	Text       MessageType = 0  // Text message
	Image      MessageType = 1  // Image message
	LargeEmoji MessageType = 2  // Large emoji message
	Emoji      MessageType = 3  // Emoji message
	File       MessageType = 4  // File message
	Card       MessageType = 5  // Card message
	InnerLink  MessageType = 6  // Inner link message
	RecallText MessageType = 16 // Recall text message
)

// GroupMemberType represents the type of group member
type GroupMemberType int

// Group member types
const (
	Member  GroupMemberType = 0
	Manager GroupMemberType = 1
	Owner   GroupMemberType = 2
)

// FileMetadata represents file metadata for upload
type FileMetadata struct {
	Size     string `json:"size"`
	GroupID  string `json:"groupid"`
	Hash     string `json:"hash"`
	Filename string `json:"filename"`
}
