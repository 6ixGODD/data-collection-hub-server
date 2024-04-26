package admin

type GetDataStatisticRequest struct {
	StartDate string `json:"start_date" validate:""`
	EndDate   string `json:"end_date" validate:""`
}

type GetUserStatisticRequest struct {
	UserUUID string `json:"user_uuid" validate:""`
}

type GetUserStatisticListRequest struct {
	Page            int    `json:"page" validate:"required"`
	Rank            string `json:"rank" validate:""`
	LastLoginBefore string `json:"last_login_before" validate:""`
	LastLoginAfter  string `json:"last_login_after" validate:""`
	CreatedBefore   string `json:"created_before" validate:""`
	CreatedAfter    string `json:"created_after" validate:""`
}

type GetInstructionDataRequest struct {
	InstructionDataID string `json:"instruction_data_id" validate:"required"`
}

type GetInstructionListRequest struct {
	Page         int    `json:"page" validate:"required"`
	UserUUID     string `json:"user_uuid" validate:""`
	UpdateBefore string `json:"update_before" validate:""`
	UpdateAfter  string `json:"update_after" validate:""`
	Theme        string `json:"theme" validate:""`
	Status       int    `json:"status" validate:""`
}

type ApproveInstructionDataRequest struct {
	InstructionDataID string `json:"instruction_data_id" validate:"required"`
}

type RejectInstructionDataRequest struct {
	InstructionDataID string `json:"instruction_data_id" validate:"required"`
}

type UpdateInstructionDataRequest struct {
	InstructionDataID string `json:"instruction_data_id" validate:"required"`
	Instruction       string `json:"instruction" validate:"required"`
	Input             string `json:"input" validate:"required"`
	Output            string `json:"output" validate:"required"`
	Theme             string `json:"theme" validate:"required"`
	Source            string `json:"source" validate:"required"`
	Note              string `json:"note" validate:""`
}

type DeleteInstructionDataRequest struct {
	InstructionDataID string `json:"instruction_data_id" validate:"required"`
}

type InsertNoticeRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	Type    string `json:"type" validate:"required"`
}

type UpdateNoticeRequest struct {
	NoticeID string `json:"notice_id" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Content  string `json:"content" validate:"required"`
	Type     string `json:"type" validate:"required"`
}

type DeleteNoticeRequest struct {
	NoticeID string `json:"notice_id" validate:"required"`
}

type InsertUserRequest struct {
	Username     string `json:"username" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
	Role         string `json:"role" validate:"required"`
	Organization string `json:"organization" validate:"required"`
}

type GetUserRequest struct {
	UserUUID string `json:"user_uuid" validate:"required"`
}

type GetUserListRequest struct {
	Page            int    `json:"page" validate:"required"`
	Role            string `json:"role" validate:""`
	LastLoginBefore string `json:"last_login_before" validate:""`
	LastLoginAfter  string `json:"last_login_after" validate:""`
	CreatedBefore   string `json:"created_before" validate:""`
	CreatedAfter    string `json:"created_after" validate:""`
}

type UpdateUserRequest struct {
	UserUUID     string `json:"user_uuid" validate:"required"`
	Username     string `json:"username" validate:""`
	Email        string `json:"email" validate:"email"`
	Role         string `json:"role" validate:""`
	Organization string `json:"organization" validate:""`
}

type DeleteUserRequest struct {
	UserUUID string `json:"user_uuid" validate:"required"`
}

type ChangeUserPasswordRequest struct {
	UserUUID    string `json:"user_uuid" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

type InsertDocumentationRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type UpdateDocumentationRequest struct {
	DocumentationID string `json:"documentation_id" validate:"required"`
	Title           string `json:"title" validate:"required"`
	Content         string `json:"content" validate:"required"`
}

type DeleteDocumentationRequest struct {
	DocumentationID string `json:"documentation_id" validate:"required"`
}

type GetLoginLogRequest struct {
	LoginLogID string `json:"login_log_id" validate:"required"`
}

type GetLoginLogListRequest struct {
	Page          int    `json:"page" validate:"required"`
	Query         string `json:"query" validate:""`
	CreatedBefore string `json:"created_before" validate:""`
	CreatedAfter  string `json:"created_after" validate:""`
}

type GetOperationLogRequest struct {
	OperationLogID string `json:"operation_log_id" validate:"required"`
}

type GetOperationLogListRequest struct {
	Page          int    `json:"page" validate:"required"`
	Query         string `json:"query" validate:""`
	Operation     string `json:"operation" validate:""`
	CreatedBefore string `json:"created_before" validate:""`
	CreatedAfter  string `json:"created_after" validate:""`
}

type GetErrorLogRequest struct {
	ErrorLogID string `json:"error_log_id" validate:"required"`
}

type GetErrorLogListRequest struct {
	Page          int    `json:"page" validate:"required"`
	RequestURL    string `json:"request_url" validate:""`
	ErrorCode     int    `json:"error_code" validate:""`
	CreatedBefore string `json:"created_before" validate:""`
	CreatedAfter  string `json:"created_after" validate:""`
}
