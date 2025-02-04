package schema

type Message struct {
	ID         string `json:"id"`
	UserID     uint   `json:"user_id"`
	TaskID     uint   `json:"task_id"`
	ProjectID  uint   `json:"project_id"`
	Content    string `json:"content"`
	UserName   string `json:"user_name"`
	UserAvatar string `json:"user_avatar"`
	UserColor  string `json:"user_color"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"update_at"`
}

type Attachment struct {
	ID        string `json:"id"`
	MessageID string `json:"message_id"`
	FileName  string `json:"file_name"`
	FileType  string `json:"file_type"`
	FileURL   string `json:"file_url"`
	FileSize  int    `json:"file_size"`
	CreatedAt string `json:"created_at"`
}
