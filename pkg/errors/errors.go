package errors

import (
	"fmt"
)

const ( // Business error code
	CodeSuccess = 200

	CodeUserNotFound  = 1001
	CodePasswordWrong = 1002
	CodeUserExist     = 1003

	CodeInvalidToken   = 2001
	CodeExpiredToken   = 2002
	CodePermissionDeny = 2003

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

	CodeServerBusy  = 5001
	CodeServerDown  = 5002
	CodeServerCrash = 5003

	CodeUnknownError = 9999
)

const ( // Default Message
	MessageSuccess = "Success"

	MessageUserNotFound  = "User not found"
	MessagePasswordWrong = "Password wrong"
	MessageUserExist     = "User exist"

	MessageInvalidToken   = "Invalid token"
	MessageExpiredToken   = "Expired token"
	MessagePermissionDeny = "Permission deny"

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

	MessageServerBusy  = "Server busy"
	MessageServerDown  = "Server down"
	MessageServerCrash = "Server crash"

	MessageUnknownError = "Unknown error"
)

type AppError struct {
	code    int    // Business error code
	status  int    // HTTP status code
	message string // Error message
}

func (e *AppError) Error() string {
	return e.message
}

func (e *AppError) Code() int {
	return e.code
}

func (e *AppError) Status() int {
	return e.status
}

func NewAppError(code, status int, message string) *AppError {
	return &AppError{
		code:    code,
		status:  status,
		message: message,
	}
}

func NewAppErrorf(code, status int, format string, a ...interface{}) *AppError {
	return &AppError{
		code:    code,
		status:  status,
		message: fmt.Sprintf(format, a...),
	}
}

func NewAppErrorWithCause(code, status int, message string, cause error) *AppError {
	return &AppError{
		code:    code,
		status:  status,
		message: fmt.Sprintf("%s: %s", message, cause.Error()),
	}
}

func UserNotFound(err error) *AppError {
	return NewAppErrorWithCause(CodeUserNotFound, 404, MessageUserNotFound, err)
}

func PasswordWrong(err error) *AppError {
	return NewAppErrorWithCause(CodePasswordWrong, 400, MessagePasswordWrong, err)
}

func UserExist(err error) *AppError {
	return NewAppErrorWithCause(CodeUserExist, 400, MessageUserExist, err)
}

func InvalidToken(err error) *AppError {
	return NewAppErrorWithCause(CodeInvalidToken, 401, MessageInvalidToken, err)
}

func ExpiredToken(err error) *AppError {
	return NewAppErrorWithCause(CodeExpiredToken, 401, MessageExpiredToken, err)
}

func PermissionDeny(err error) *AppError {
	return NewAppErrorWithCause(CodePermissionDeny, 403, MessagePermissionDeny, err)
}

func InvalidRequest(err error) *AppError {
	return NewAppErrorWithCause(CodeInvalidRequest, 400, MessageInvalidRequest, err)
}

func InvalidParams(err error) *AppError {
	return NewAppErrorWithCause(CodeInvalidParams, 400, MessageInvalidParams, err)
}

func MissingParams(err error) *AppError {
	return NewAppErrorWithCause(CodeMissingParams, 400, MessageMissingParams, err)
}

func ConnError(err error) *AppError {
	return NewAppErrorWithCause(CodeConnError, 500, MessageConnError, err)
}

func ReadError(err error) *AppError {
	return NewAppErrorWithCause(CodeReadError, 500, MessageReadError, err)
}

func WriteError(err error) *AppError {
	return NewAppErrorWithCause(CodeWriteError, 500, MessageWriteError, err)
}

func DBError(err error) *AppError {
	return NewAppErrorWithCause(CodeDBError, 500, MessageDBError, err)
}

func CacheError(err error) *AppError {
	return NewAppErrorWithCause(CodeCacheError, 500, MessageCacheError, err)
}

func MongoError(err error) *AppError {
	return NewAppErrorWithCause(CodeMongoError, 500, MessageMongoError, err)
}

func RedisError(err error) *AppError {
	return NewAppErrorWithCause(CodeRedisError, 500, MessageRedisError, err)
}

func ServerBusy(err error) *AppError {
	return NewAppErrorWithCause(CodeServerBusy, 503, MessageServerBusy, err)
}

func ServerDown(err error) *AppError {
	return NewAppErrorWithCause(CodeServerDown, 503, MessageServerDown, err)
}

func ServerCrash(err error) *AppError {
	return NewAppErrorWithCause(CodeServerCrash, 503, MessageServerCrash, err)
}

func UnknownError(err error) *AppError {
	return NewAppErrorWithCause(CodeUnknownError, 500, MessageUnknownError, err)
}
