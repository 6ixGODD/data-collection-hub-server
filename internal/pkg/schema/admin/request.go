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
		Page               *int64  `query:"page" validate:"required,numeric,min=1"`
		PageSize           *int64  `query:"page_size" validate:"required,numeric,min=1,max=100"`
		LastLoginStartTime *string `query:"last_login_start_time" validate:"datetime"`
		LastLoginEndTime   *string `query:"last_login_end_time" validate:"datetime"`
		CreateStartTime    *string `query:"create_start_time" validate:"datetime"`
		CreateEndTime      *string `query:"create_end_time" validate:"datetime"`
	}

	GetInstructionDataRequest struct {
		InstructionDataID *string `query:"instruction_data_id" validate:"required,mongodb"`
	}

	GetInstructionDataListRequest struct {
		Page            *int64  `query:"page" validate:"required,numeric,min=1"`
		PageSize        *int64  `query:"page_size" validate:"required,numeric,min=1,max=100"`
		Desc            *bool   `query:"desc" validate:""`
		UserID          *string `query:"user_id" validate:"mongodb"`
		CreateStartTime *string `query:"create_start_time" validate:"datetime"`
		CreateEndTime   *string `query:"create_end_time" validate:"datetime"`
		UpdateStartTime *string `query:"update_start_time" validate:"datetime"`
		UpdateEndTime   *string `query:"update_end_time" validate:"datetime"`
		Theme           *string `query:"theme" validate:""`
		Status          *string `query:"status" validate:""`
		Query           *string `query:"query" validate:""`
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
		UserID            *string `json:"user_id" validate:""`
		Instruction       *string `json:"instruction" validate:""`
		Input             *string `json:"input" validate:""`
		Output            *string `json:"output" validate:""`
		Theme             *string `json:"theme" validate:""`
		Source            *string `json:"source" validate:""`
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
		Page               *int64  `query:"page" validate:"required,numeric,min=1"`
		PageSize           *int64  `query:"page_size" validate:"required,numeric,min=1,max=100"`
		Desc               *bool   `query:"desc" validate:""`
		Role               *string `query:"role" validate:""`
		LastLoginTimeStart *string `query:"last_login_start_time" validate:"datetime"`
		LastLoginTimeEnd   *string `query:"last_login_end_time" validate:"datetime"`
		CreateTimeStart    *string `query:"create_start_time" validate:"datetime"`
		CreateTimeEnd      *string `query:"create_end_time" validate:"datetime"`
		Query              *string `query:"query" validate:""`
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
		Page            *int64  `query:"page" validate:"required,numeric,min=1"`
		PageSize        *int64  `query:"page_size" validate:"required,numeric,min=1,max=100"`
		Desc            *bool   `query:"desc" validate:""`
		Query           *string `query:"query" validate:""`
		CreateStartTime *string `query:"create_start_time" validate:"datetime"`
		CreateEndTime   *string `query:"create_end_time" validate:"datetime"`
	}

	GetOperationLogRequest struct {
		OperationLogID *string `query:"operation_log_id" validate:"required,mongodb"`
	}

	GetOperationLogListRequest struct {
		Page            *int64  `query:"page" validate:"required,numeric,min=1"`
		PageSize        *int64  `query:"page_size" validate:"required,numeric,min=1,max=100"`
		Desc            *bool   `query:"desc" validate:""`
		Query           *string `query:"query" validate:""`
		Operation       *string `query:"operation" validate:""`
		CreateStartTime *string `query:"create_start_time" validate:"datetime"`
		CreateEndTime   *string `query:"create_end_time" validate:"datetime"`
	}

	GetErrorLogRequest struct {
		ErrorLogID *string `query:"error_log_id" validate:"required,mongodb"`
	}

	GetErrorLogListRequest struct {
		Page            *int64  `query:"page" validate:"required,numeric"`
		PageSize        *int64  `query:"page_size" validate:"required,numeric,min=1,max=100"`
		Desc            *bool   `query:"desc" validate:""`
		RequestURL      *string `query:"request_url" validate:""`
		ErrorCode       *string `query:"error_code" validate:""`
		CreateStartTime *string `query:"create_start_time" validate:"datetime"`
		CreateEndTime   *string `query:"create_end_time" validate:"datetime"`
	}
)
