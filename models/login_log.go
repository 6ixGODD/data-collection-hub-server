package models

// LoginLogModel Login Log Model
type LoginLogModel struct {
	LoginLogID string `json:"_id"`        // MongoDB ObjectId
	UserUUID   string `json:"user_uuid"`  // User UUID
	Username   string `json:"username"`   // Username (for space-time trade-off)
	Email      string `json:"email"`      // Email (for space-time trade-off)
	IPAddress  string `json:"ip_address"` // IP Address
	UserAgent  string `json:"user_agent"` // User Agent
	CreatedAt  string `json:"created_at"` // Created Time in ISO 8601
}
