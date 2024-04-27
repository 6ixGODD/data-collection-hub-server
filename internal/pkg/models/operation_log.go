package models

type OperationLogModel struct {
	OperationLogID string `json:"operation_log_id" bson:"_id"`    // MongoDB ObjectId
	UserUUID       string `json:"user_uuid" bson:"user_uuid"`     // User UUID
	Username       string `json:"username" bson:"username"`       // Username (for space-time trade-off)
	Email          string `json:"email" bson:"email"`             // Email (for space-time trade-off)
	IPAddress      string `json:"ip_address" bson:"ip_address"`   // IP Address
	UserAgent      string `json:"user_agent" bson:"user_agent"`   // User Agent
	Operation      string `json:"operation" bson:"operation"`     // Operation, 'CREATE' | 'UPDATE' | 'DELETE'
	EntityUUID     string `json:"entity_uuid" bson:"entity_uuid"` // Entity(Instruction Data) UUID
	CreatedAt      string `json:"created_at" bson:"created_at"`   // Created Time in ISO 8601
}
