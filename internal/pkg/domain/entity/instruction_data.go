package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InstructionDataModel struct {
	InstructionDataID primitive.ObjectID `json:"instruction_data_id" bson:"_id"` // Mongo ObjectID
	UserID            primitive.ObjectID `json:"user_id" bson:"user_id"`         // User ID
	Username          string             `json:"username" bson:"username"`       // Username (for space-time trade-off)
	Row               struct {           // Row Data in alpaca format
		Instruction string `json:"instruction" bson:"instruction"` // Instruction
		Input       string `json:"input" bson:"input"`             // Input
		Output      string `json:"output" bson:"output"`           // Output
	} `json:"row" bson:"row"`
	Theme  string   `json:"theme" bson:"theme"`   // Theme
	Source string   `json:"source" bson:"source"` // Source
	Note   string   `json:"note" bson:"note"`     // Note (Optional)
	Status struct { // Status
		Code    string `json:"code" bson:"code"`       // Status Code, 'PENDING' | 'APPROVED' | 'REJECTED'
		Message string `json:"message" bson:"message"` // Status Error
	} `json:"status" bson:"status"`
	Deleted   bool      `json:"deleted" bson:"deleted"`       // Deleted Flag
	CreatedAt time.Time `json:"created_at" bson:"created_at"` // Created Time in ISO 8601
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"` // Updated Time in ISO 8601
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"` // Deleted Time in ISO 8601
}
