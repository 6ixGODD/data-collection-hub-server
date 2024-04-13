package models

// InstructionDataModel Instruction Data Model
type InstructionDataModel struct {
	InstructionDataID string   `json:"_id"`       // MongoDB ObjectId
	UserUUID          string   `json:"user_uuid"` // User UUID
	Username          string   `json:"username"`  // Username (for space-time trade-off)
	RowUUID           string   `json:"row_uuid"`  // Row Data Index (UUID)
	Row               struct { // Row Data
		Instruction string `json:"instruction"`
		Input       string `json:"input"`
		Output      string `json:"output"`
	} `json:"row"`
	Theme  string   `json:"theme"`  // Theme
	Source string   `json:"source"` // Source
	Note   string   `json:"note"`   // Note (Optional)
	Status struct { // Status
		Code    string `json:"code"`    // Status Code, 'PENDING' | 'APPROVED' | 'REJECTED'
		Message string `json:"message"` // Status Message
	} `json:"status"`
	Deleted   bool   `json:"deleted"`    // Deleted Flag
	CreatedAt string `json:"created_at"` // Created Time in ISO 8601
	UpdatedAt string `json:"updated_at"` // Updated Time in ISO 8601
}
