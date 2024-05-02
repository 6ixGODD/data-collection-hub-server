package models

type NoticeModel struct {
	NoticeID  string `json:"notice_id" bson:"_id"`         // MongoClient ObjectId
	Title     string `json:"title" bson:"title"`           // Title
	Content   string `json:"content" bson:"content"`       // Content in Markdown format
	Type      string `json:"type" bson:"type"`             // Type, 'URGENT' | 'NORMAL'
	CreatedAt string `json:"created_at" bson:"created_at"` // Created Time in ISO 8601
	UpdatedAt string `json:"updated_at" bson:"updated_at"` // Updated Time in ISO 8601
}
