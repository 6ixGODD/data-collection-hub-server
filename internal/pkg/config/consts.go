package config

const ( // Context keys
	KeyUserID    = "UserID"
	KeyRole      = "Role"
	KeyRequestID = "RequestID"
)

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

const (
	DocumentationCollectionName   = "documentation"
	ErrorLogCollectionName        = "error_log"
	NoticeCollectionName          = "notice"
	InstructionDataCollectionName = "instruction_data"
	LoginLogCollectionName        = "login_log"
	OperationLogCollectionName    = "operation_log"
	UserCollectionName            = "user"
)

const (
	NoticeCachePrefix        = "dao:notice"
	UserCachePrefix          = "dao:user"
	DocumentationCachePrefix = "dao:documentation"
)
