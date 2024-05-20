package errors

import (
	"github.com/gofiber/fiber/v2"
)

func UserNotFound(err error) *AppError {
	return NewAppErrorWithCause(CodeUserNotFound, fiber.StatusNotFound, MessageUserNotFound, err)
}

func PasswordWrong(err error) *AppError {
	return NewAppErrorWithCause(CodePasswordWrong, fiber.StatusBadRequest, MessagePasswordWrong, err)
}

func UserExist(err error) *AppError {
	return NewAppErrorWithCause(CodeUserExist, fiber.StatusBadRequest, MessageUserExist, err)
}

func InvalidToken(err error) *AppError {
	return NewAppErrorWithCause(CodeInvalidToken, fiber.StatusUnauthorized, MessageInvalidToken, err)
}

func ExpiredToken(err error) *AppError {
	return NewAppErrorWithCause(CodeExpiredToken, fiber.StatusUnauthorized, MessageExpiredToken, err)
}

func TokenMissed(err error) *AppError {
	return NewAppErrorWithCause(CodeTokenMissed, fiber.StatusUnauthorized, MessageTokenMissed, err)
}

func TokenGenerateFailed(err error) *AppError {
	return NewAppErrorWithCause(
		CodeTokenGenerateFailed, fiber.StatusInternalServerError, MessageTokenGenerateFailed, err,
	)
}

func PermissionDeny(err error) *AppError {
	return NewAppErrorWithCause(CodePermissionDeny, fiber.StatusForbidden, MessagePermissionDeny, err)
}

func InvalidRequest(err error) *AppError {
	return NewAppErrorWithCause(CodeInvalidRequest, fiber.StatusBadRequest, MessageInvalidRequest, err)
}

func InvalidParams(err error) *AppError {
	return NewAppErrorWithCause(CodeInvalidParams, fiber.StatusBadRequest, MessageInvalidParams, err)
}

func MissingParams(err error) *AppError {
	return NewAppErrorWithCause(CodeMissingParams, fiber.StatusBadRequest, MessageMissingParams, err)
}

func ConnError(err error) *AppError {
	return NewAppErrorWithCause(CodeConnError, fiber.StatusInternalServerError, MessageConnError, err)
}

func ReadError(err error) *AppError {
	return NewAppErrorWithCause(CodeDBReadError, fiber.StatusInternalServerError, MessageReadError, err)
}

func WriteError(err error) *AppError {
	return NewAppErrorWithCause(CodeDBWriteError, fiber.StatusInternalServerError, MessageWriteError, err)
}

func DBError(err error) *AppError {
	return NewAppErrorWithCause(CodeDBError, fiber.StatusInternalServerError, MessageMongoError, err)
}

func CacheError(err error) *AppError {
	return NewAppErrorWithCause(CodeCacheError, fiber.StatusInternalServerError, MessageRedisError, err)
}

func ServerBusy(err error) *AppError {
	return NewAppErrorWithCause(CodeServerBusy, fiber.StatusServiceUnavailable, MessageServerBusy, err)
}

func ServerDown(err error) *AppError {
	return NewAppErrorWithCause(CodeServerDown, fiber.StatusServiceUnavailable, MessageServerDown, err)
}

func ServerCrash(err error) *AppError {
	return NewAppErrorWithCause(CodeServerCrash, fiber.StatusServiceUnavailable, MessageServerCrash, err)
}

func ServiceError(err error) *AppError {
	return NewAppErrorWithCause(CodeServiceError, fiber.StatusInternalServerError, MessageServiceError, err)
}

func UnknownError(err error) *AppError {
	return NewAppErrorWithCause(CodeUnknownError, fiber.StatusInternalServerError, MessageUnknownError, err)
}
