package config

import (
	"data-collection-hub-server/pkg/zap"
)

// Context Key
const (
	UserIDKey    = zap.UserIDKey
	KeyRole      = "Role"
	KeyRequestID = zap.RequestIDKey
)

// Enum Values
const (
	InstructionDataStatusPending  = "PENDING"
	InstructionDataStatusApproved = "APPROVED"
	InstructionDataStatusRejected = "REJECTED"

	NoticeTypeUrgent = "URGENT"
	NoticeTypeNormal = "NORMAL"

	OperationTypeCreate = "CREATE"
	OperationTypeUpdate = "UPDATE"
	OperationTypeDelete = "DELETE"

	EntityTypeInstruction   = "INSTRUCTION"
	EntityTypeUser          = "USER"
	EntityTypeDocumentation = "DOCUMENTATION"
	EntityTypeNotice        = "NOTICE"

	OperationStatusSuccess = "SUCCESS"
	OperationStatusFailure = "FAILURE"

	UserRoleUser  = "USER"
	UserRoleAdmin = "ADMIN"
)

// MongoDB Collection Name
const (
	DocumentationCollectionName   = "documentation"
	NoticeCollectionName          = "notice"
	InstructionDataCollectionName = "instruction_data"
	LoginLogCollectionName        = "login_log"
	OperationLogCollectionName    = "operation_log"
	UserCollectionName            = "user"
)

// Cache Prefix / Key
const (
	NoticeCachePrefix        = "dao:notice"
	UserCachePrefix          = "dao:user"
	DocumentationCachePrefix = "dao:documentation"

	LoginLogCacheKey     = "log:login"
	OperationLogCacheKey = "log:operation"
)
