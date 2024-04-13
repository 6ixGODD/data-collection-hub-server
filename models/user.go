package models

// UserModel User Model
type UserModel struct {
	ID           string `json:"_id"`          // MongoDB ObjectId
	UUID         string `json:"uuid"`         // User UUID
	Username     string `json:"username"`     // Username
	Email        string `json:"email"`        // Email
	Password     string `json:"password"`     // Password crypt
	Role         string `json:"role"`         // Role, 'USER' | 'ADMIN'
	Organization string `json:"organization"` // Organization
	LastLogin    string `json:"last_login"`   // Last Login Time in ISO 8601
	Deleted      bool   `json:"deleted"`      // Deleted Flag
	CreatedAt    string `json:"created_at"`   // Created Time in ISO 8601
	UpdatedAt    string `json:"updated_at"`   // Updated Time in ISO 8601
	DeletedAt    string `json:"deleted_at"`   // Deleted Time in ISO 8601
}
