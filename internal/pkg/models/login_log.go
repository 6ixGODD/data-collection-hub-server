package models

type LoginLogModel struct {
	LoginLogID string `json:"login_log_id" bson:"_id"`      // Mongo ObjectId
	UserUUID   string `json:"user_uuid" bson:"user_uuid"`   // User UUID
	Username   string `json:"username" bson:"username"`     // Username (for space-time trade-off)
	Email      string `json:"email" bson:"email"`           // Email (for space-time trade-off)
	IPAddress  string `json:"ip_address" bson:"ip_address"` // IP Address
	UserAgent  string `json:"user_agent" bson:"user_agent"` // User Agent
	CreatedAt  string `json:"created_at" bson:"created_at"` // Created Time in ISO 8601
}
