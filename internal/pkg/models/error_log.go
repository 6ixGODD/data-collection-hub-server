package models

type ErrorLogModel struct {
	ErrorLogID     string `json:"error_log_id" bson:"_id"`                // MongoClient ObjectId
	UserUUID       string `json:"user_uuid" bson:"user_uuid"`             // User UUID
	Username       string `json:"username" bson:"username"`               // Username (for space-time trade-off)
	IPAddress      string `json:"ip_address" bson:"ip_address"`           // IP Address
	UserAgent      string `json:"user_agent" bson:"user_agent"`           // User Agent
	RequestURL     string `json:"request_url" bson:"request_url"`         // Request URL
	RequestMethod  string `json:"request_method" bson:"request_method"`   // Request Method
	RequestPayload string `json:"request_payload" bson:"request_payload"` // Request Payload
	ErrorCode      string `json:"error_code" bson:"error_code"`           // Error Code, e.g. '500', '502'
	ErrorMsg       string `json:"error_msg" bson:"error_msg"`             // Error Message
	Stack          string `json:"stack" bson:"stack"`                     // Error Stack Trace
	CreatedAt      string `json:"created_at" bson:"created_at"`           // Created Time in ISO 8601
}
