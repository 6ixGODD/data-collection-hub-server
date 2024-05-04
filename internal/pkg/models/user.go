package models

type UserModel struct {
	UserID       string `json:"user_id" bson:"_id"`               // Mongo ObjectId
	UUID         string `json:"uuid" bson:"uuid"`                 // User UUID
	Username     string `json:"username" bson:"username"`         // Username
	Email        string `json:"email" bson:"email"`               // Email
	Password     string `json:"password" bson:"password"`         // Password crypt
	Role         string `json:"role" bson:"role"`                 // Role, 'USER' | 'ADMIN'
	Organization string `json:"organization" bson:"organization"` // Organization
	LastLogin    string `json:"last_login" bson:"last_login"`     // Last Login Time in ISO 8601
	Deleted      bool   `json:"deleted" bson:"deleted"`           // Deleted Flag
	CreatedAt    string `json:"created_at" bson:"created_at"`     // Created Time in ISO 8601
	UpdatedAt    string `json:"updated_at" bson:"updated_at"`     // Updated Time in ISO 8601
	DeletedAt    string `json:"deleted_at" bson:"deleted_at"`     // Deleted Time in ISO 8601
}
