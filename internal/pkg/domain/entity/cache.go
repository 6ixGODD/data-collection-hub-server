package entity

import (
	"time"
)

type CacheList struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

type LoginLogCache struct {
	UserIDHex string    `json:"user_id_hex"` // User ID in Hex
	Username  string    `json:"username"`    // Username (for space-time trade-off)
	Email     string    `json:"email"`       // Email (for space-time trade-off)
	IPAddress string    `json:"ip_address"`  // IP Address
	UserAgent string    `json:"user_agent"`  // User Agent
	CreatedAt time.Time `json:"created_at"`  // Created Time in ISO 8601
}

type OperationLogCache struct {
	UserIDHex   string    `json:"user_id_hex"`   // User ID in Hex
	Username    string    `json:"username"`      // Username (for space-time trade-off)
	Email       string    `json:"email"`         // Email (for space-time trade-off)
	IPAddress   string    `json:"ip_address"`    // IP Address
	UserAgent   string    `json:"user_agent"`    // User Agent
	Operation   string    `json:"operation"`     // Operation, 'CREATE' | 'UPDATE' | 'DELETE'
	EntityIDHex string    `json:"entity_id_hex"` // Entity(Instruction Data) ID in Hex
	EntityType  string    `json:"entity_type"`   // Entity NoticeType, 'INSTRUCTION' | 'USER' | 'DOCUMENTATION' | 'NOTICE'
	Description string    `json:"description"`   // Description of Operation
	Status      string    `json:"status"`        // Status, 'SUCCESS' | 'FAILURE'
	CreatedAt   time.Time `json:"created_at"`    // Created Time in ISO 8601
}
