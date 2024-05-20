package errors

const ( // Business error code
	CodeSuccess = 200

	CodeUserNotFound  = 1001
	CodePasswordWrong = 1002
	CodeUserExist     = 1003

	CodeInvalidToken        = 2001
	CodeExpiredToken        = 2002
	CodePermissionDeny      = 2003
	CodeTokenMissed         = 2004
	CodeTokenGenerateFailed = 2005

	CodeInvalidRequest = 3001
	CodeInvalidParams  = 3002
	CodeMissingParams  = 3003

	CodeConnError  = 4001
	CodeReadError  = 4002
	CodeWriteError = 4003
	CodeDBError    = 4004
	CodeCacheError = 4005
	CodeMongoError = 4006
	CodeRedisError = 4007

	CodeServerBusy   = 5001
	CodeServerDown   = 5002
	CodeServerCrash  = 5003
	CodeServiceError = 5004

	CodeUnknownError = 9999
)

const ( // Default Message
	MessageSuccess = "Success"

	MessageUserNotFound  = "User not found"
	MessagePasswordWrong = "Password wrong"
	MessageUserExist     = "User exist"

	MessageInvalidToken        = "Invalid token"
	MessageExpiredToken        = "Expired token"
	MessagePermissionDeny      = "Permission deny"
	MessageTokenMissed         = "Token missed"
	MessageTokenGenerateFailed = "Token generate failed"

	MessageInvalidRequest = "Invalid request"
	MessageInvalidParams  = "Invalid params"
	MessageMissingParams  = "Missing params"

	MessageConnError  = "Connection error"
	MessageReadError  = "Read error"
	MessageWriteError = "Write error"
	MessageDBError    = "Database error"
	MessageCacheError = "Cache error"
	MessageMongoError = "Mongo error"
	MessageRedisError = "Redis error"

	MessageServerBusy   = "Server busy"
	MessageServerDown   = "Server down"
	MessageServerCrash  = "Server crash"
	MessageServiceError = "core error"

	MessageUnknownError = "Unknown error"
)
