package models

type InstructionDataModel struct {
	InstructionDataID string   `json:"_id" bson:"_id"`             // MongoDB ObjectID
	UserUUID          string   `json:"user_uuid" bson:"user_uuid"` // User UUID
	Username          string   `json:"username" bson:"username"`   // Username (for space-time trade-off)
	RowUUID           string   `json:"row_uuid" bson:"row_uuid"`   // Row Data Index (UUID)
	Row               struct { // Row Data in alpaca
		Instruction string `json:"instruction" bson:"instruction"` // Instruction
		Input       string `json:"input" bson:"input"`             // Input
		Output      string `json:"output" bson:"output"`           // Output
	} `json:"row" bson:"row"`
	Theme  string   `json:"theme" bson:"theme"`   // Theme
	Source string   `json:"source" bson:"source"` // Source
	Note   string   `json:"note" bson:"note"`     // Note (Optional)
	Status struct { // Status
		Code    string `json:"code" bson:"code"`       // Status Code, 'PENDING' | 'APPROVED' | 'REJECTED'
		Message string `json:"message" bson:"message"` // Status Message
	} `json:"status" bson:"status"`
	Deleted   bool   `json:"deleted" bson:"deleted"`       // Deleted Flag
	CreatedAt string `json:"created_at" bson:"created_at"` // Created Time in ISO 8601
	UpdatedAt string `json:"updated_at" bson:"updated_at"` // Updated Time in ISO 8601
	DeletedAt string `json:"deleted_at" bson:"deleted_at"` // Deleted Time in ISO 8601
}
