package models

// DocumentationModel Documentation Model
type DocumentationModel struct {
	DocumentID string `json:"_id"`        // MongoDB ObjectId
	Title      string `json:"title"`      // Title of the document
	Content    string `json:"content"`    // Content of the document
	CreatedAt  string `json:"created_at"` // Create Time in ISO 8601
	UpdatedAt  string `json:"updated_at"` // Update Time in ISO 8601
}
