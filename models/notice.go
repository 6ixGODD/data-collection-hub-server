package models

// NoticeModel Notice Model
type NoticeModel struct {
	NoticeID  string `json:"_id"`        // MongoDB ObjectId
	Title     string `json:"title"`      // Title
	Content   string `json:"content"`    // Content in Markdown format
	Type      string `json:"type"`       // Type, 'URGENT' | 'NORMAL'
	CreatedAt string `json:"created_at"` // Created Time in ISO 8601
	UpdatedAt string `json:"updated_at"` // Updated Time in ISO 8601
}
