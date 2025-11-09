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

const (
	Text        MessageType = 0
	Image       MessageType = 1
	LargeEmoji  MessageType = 2
	Emoji       MessageType = 3
	File        MessageType = 4
	Card        MessageType = 5
	InnerLink   MessageType = 6
	Recall_Text MessageType = 16
)

// GroupMemberType represents the type of group member
type GroupMemberType int

const (
	Member  GroupMemberType = 0
	Manager GroupMemberType = 1
	Owner   GroupMemberType = 2
)

// FileMetadata represents file metadata for upload
type FileMetadata struct {
	Size     int64  `json:"size"`
	GroupID  int64  `json:"groupid"`
	Hash     string `json:"hash"`
	Filename string `json:"filename"`
}