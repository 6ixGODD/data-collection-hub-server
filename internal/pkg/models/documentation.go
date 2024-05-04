package models

type DocumentationModel struct {
	DocumentID string `json:"document_id" bson:"_id"`       // Mongo ObjectId
	Title      string `json:"title" bson:"title"`           // Title of the document
	Content    string `json:"content" bson:"content"`       // Content of the document
	CreatedAt  string `json:"created_at" bson:"created_at"` // Create Time in ISO 8601
	UpdatedAt  string `json:"updated_at" bson:"updated_at"` // Update Time in ISO 8601
}
