package models

// ErrorLog Error Log Model
type ErrorLog struct {
	ErrorLogID     string `json:"_id"`             // MongoDB ObjectId
	UserUUID       string `json:"user_uuid"`       // User UUID
	Username       string `json:"username"`        // Username (for space-time trade-off)
	IPAddress      string `json:"ip_address"`      // IP Address
	UserAgent      string `json:"user_agent"`      // User Agent
	RequestURL     string `json:"request_url"`     // Request URL
	RequestMethod  string `json:"request_method"`  // Request Method
	RequestPayload string `json:"request_payload"` // Request Payload
	ErrorCode      string `json:"error_code"`      // Error Code, e.g. '500', '502'
	ErrorMsg       string `json:"error_msg"`       // Error Message
	Stack          string `json:"stack"`           // Error Stack Trace
	CreatedAt      string `json:"created_at"`      // Created Time in ISO 8601
}
