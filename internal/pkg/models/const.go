package models

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
