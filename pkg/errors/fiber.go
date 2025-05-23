package errors

import (
	"github.com/gofiber/fiber/v2"
)

func NotAuthorized(err error) *AppError {
	return NewAppErrorWithCause(CodeNotAuthorized, fiber.StatusUnauthorized, "Not authorized", err)
}

func AuthFailed(err error) *AppError {
	return NewAppErrorWithCause(CodeAuthFailed, fiber.StatusUnauthorized, "Authentication failed", err)
}

func TokenInvalid(err error) *AppError {
	return NewAppErrorWithCause(CodeTokenInvalid, fiber.StatusUnauthorized, "Token invalid", err)
}

func TokenExpired(err error) *AppError {
	return NewAppErrorWithCause(CodeTokenExpired, fiber.StatusUnauthorized, "Token expired", err)
}

func TokenMissed(err error) *AppError {
	return NewAppErrorWithCause(CodeTokenMissed, fiber.StatusUnauthorized, "Token missed", err)
}

func PermissionDeny(err error) *AppError {
	return NewAppErrorWithCause(CodePermissionDeny, fiber.StatusForbidden, "Permission deny", err)
}

func InvalidRequest(err error) *AppError {
	return NewAppErrorWithCause(CodeInvalidRequest, fiber.StatusBadRequest, "Invalid request", err)
}

func Idempotency(err error) *AppError {
	return NewAppErrorWithCause(CodeIdempotency, fiber.StatusBadRequest, "Idempotency check failed", err)
}

func NotFound(err error) *AppError {
	return NewAppErrorWithCause(CodeNotFound, fiber.StatusNotFound, "Not found", err)
}

func OperationFailed(err error) *AppError {
	return NewAppErrorWithCause(CodeOperationFailed, fiber.StatusInternalServerError, "Operation failed", err)
}

func DuplicateKeyError(err error) *AppError {
	return NewAppErrorWithCause(CodeDuplicateKey, fiber.StatusBadRequest, "Duplicate key error", err)
}

func ServerBusy(err error) *AppError {
	return NewAppErrorWithCause(CodeServerBusy, fiber.StatusInternalServerError, "Server busy", err)
}

func ServiceError(err error) *AppError {
	return NewAppErrorWithCause(CodeServiceError, fiber.StatusInternalServerError, "Service error", err)
}
