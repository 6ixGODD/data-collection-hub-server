package admin

type (
	GetDataStatisticRequest struct {
		StartDate *string `query:"start_date" validate:"datetime"`
		EndDate   *string `query:"end_date" validate:"datetime"`
	}

	GetUserStatisticRequest struct {
		UserID *string `query:"user_id" validate:"mongodb"`
	}

	GetUserStatisticListRequest struct {
		Page            *int    `query:"page" validate:"required,numeric"`
		LastLoginBefore *string `query:"last_login_before" validate:"datetime"`
		LastLoginAfter  *string `query:"last_login_after" validate:"datetime"`
		CreatedBefore   *string `query:"created_before" validate:"datetime"`
		CreatedAfter    *string `query:"created_after" validate:"datetime"`
	}

	GetInstructionDataRequest struct {
		InstructionDataID *string `query:"instruction_data_id" validate:"required,mongodb"`
	}

	GetInstructionDataListRequest struct {
		Page         *int    `query:"page" validate:"required,numeric"`
		Desc         *bool   `query:"desc" validate:""`
		UserID       *string `query:"user_id" validate:"mongodb"`
		UpdateBefore *string `query:"update_before" validate:"datetime"`
		UpdateAfter  *string `query:"update_after" validate:"datetime"`
		Theme        *string `query:"theme" validate:""`
		Status       *string `query:"status" validate:""`
	}

	ApproveInstructionDataRequest struct {
		InstructionDataID *string `json:"instruction_data_id" validate:"required"`
	}

	RejectInstructionDataRequest struct {
		InstructionDataID *string `json:"instruction_data_id" validate:"required"`
		Message           *string `json:"message" validate:"required"`
	}

	UpdateInstructionDataRequest struct {
		InstructionDataID *string `json:"instruction_data_id" validate:"required"`
		Instruction       *string `json:"instruction" validate:"required"`
		Input             *string `json:"input" validate:"required"`
		Output            *string `json:"output" validate:"required"`
		Theme             *string `json:"theme" validate:"required"`
		Source            *string `json:"source" validate:"required"`
		Note              *string `json:"note" validate:""`
	}

	DeleteInstructionDataRequest struct {
		InstructionDataID *string `query:"instruction_data_id" validate:"required"`
	}

	InsertNoticeRequest struct {
		Title      *string `json:"title" validate:"required"`
		Content    *string `json:"content" validate:"required"`
		NoticeType *string `json:"notice_type" validate:"required"`
	}

	UpdateNoticeRequest struct {
		NoticeID   *string `json:"notice_id" validate:"required"`
		Title      *string `json:"title" validate:"required"`
		Content    *string `json:"content" validate:"required"`
		NoticeType *string `json:"notice_type" validate:"required"`
	}

	DeleteNoticeRequest struct {
		NoticeID *string `query:"notice_id" validate:"required"`
	}

	InsertUserRequest struct {
		Username     *string `json:"username" validate:"required"`
		Email        *string `json:"email" validate:"required,email"`
		Password     *string `json:"password" validate:"required,min=8,max=20"`
		Role         *string `json:"role" validate:"required"`
		Organization *string `json:"organization" validate:"required"`
	}

	GetUserRequest struct {
		UserID *string `query:"user_id" validate:"required,mongodb"`
	}

	GetUserListRequest struct {
		Page            *int    `query:"page" validate:"required"`
		Role            *string `query:"role" validate:""`
		LastLoginBefore *string `query:"last_login_before" validate:"datetime"`
		LastLoginAfter  *string `query:"last_login_after" validate:"datetime"`
		CreatedBefore   *string `query:"created_before" validate:"datetime"`
		CreatedAfter    *string `query:"created_after" validate:"datetime"`
	}

	UpdateUserRequest struct {
		UserID       *string `json:"user_id" validate:"required"`
		Username     *string `json:"username" validate:"omitempty,min=3,max=20"`
		Email        *string `json:"email" validate:"omitempty,email"`
		Role         *string `json:"role" validate:""`
		Organization *string `json:"organization" validate:""`
	}

	DeleteUserRequest struct {
		UserID *string `json:"user_id" validate:"required,mongodb"`
	}

	ChangeUserPasswordRequest struct {
		UserID      *string `json:"user_id" validate:"required,mongodb"`
		NewPassword *string `json:"new_password" validate:"required,min=8,max=20"`
	}

	InsertDocumentationRequest struct {
		Title   *string `json:"title" validate:"required"`
		Content *string `json:"content" validate:"required"`
	}

	UpdateDocumentationRequest struct {
		DocumentationID *string `json:"documentation_id" validate:"required,mongodb"`
		Title           *string `json:"title" validate:"required"`
		Content         *string `json:"content" validate:"required"`
	}

	DeleteDocumentationRequest struct {
		DocumentationID *string `json:"documentation_id" validate:"required,mongodb"`
	}

	GetLoginLogRequest struct {
		LoginLogID *string `query:"login_log_id" validate:"required,mongodb"`
	}

	GetLoginLogListRequest struct {
		Page          *int    `query:"page" validate:"required,numeric"`
		Query         *string `query:"query" validate:""`
		CreatedBefore *string `query:"created_before" validate:"datetime"`
		CreatedAfter  *string `query:"created_after" validate:"datetime"`
	}

	GetOperationLogRequest struct {
		OperationLogID *string `query:"operation_log_id" validate:"required,mongodb"`
	}

	GetOperationLogListRequest struct {
		Page          *int    `query:"page" validate:"required,numeric"`
		Query         *string `query:"query" validate:""`
		Operation     *string `query:"operation" validate:""`
		CreatedBefore *string `query:"created_before" validate:"datetime"`
		CreatedAfter  *string `query:"created_after" validate:"datetime"`
	}

	GetErrorLogRequest struct {
		ErrorLogID *string `query:"error_log_id" validate:"required,mongodb"`
	}

	GetErrorLogListRequest struct {
		Page          *int    `query:"page" validate:"required,numeric"`
		RequestURL    *string `query:"request_url" validate:""`
		ErrorCode     *int    `query:"error_code" validate:""`
		CreatedBefore *string `query:"created_before" validate:"datetime"`
		CreatedAfter  *string `query:"created_after" validate:"datetime"`
	}
)
