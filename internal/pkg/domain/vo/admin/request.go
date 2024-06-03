package admin

type (
	GetDataStatisticRequest struct {
		StartDate *string `query:"startDate" validate:"omitnil,rfc3339,earlierThan=EndDate"`
		EndDate   *string `query:"endDate" validate:"omitnil,rfc3339"`
	}

	GetUserStatisticRequest struct {
		UserID *string `query:"userID" validate:"required,mongodb"`
	}

	GetUserStatisticListRequest struct {
		Page               *int64  `query:"page" validate:"required,numeric,min=1"`
		PageSize           *int64  `query:"pageSize" validate:"required,numeric,min=1,max=100"`
		LastLoginStartTime *string `query:"lastLoginStartTime" validate:"omitnil,rfc3339,earlierThan=LastLoginEndTime"`
		LastLoginEndTime   *string `query:"lastLoginEndTime" validate:"omitnil,rfc3339"`
		CreateStartTime    *string `query:"createStartTime" validate:"omitnil,rfc3339,earlierThan=CreateEndTime"`
		CreateEndTime      *string `query:"createEndTime" validate:"omitnil,rfc3339"`
	}

	GetInstructionDataRequest struct {
		InstructionDataID *string `query:"instructionDataID" validate:"required,mongodb"`
	}

	GetInstructionDataListRequest struct {
		Page            *int64  `query:"page" validate:"required,numeric,min=1"`
		PageSize        *int64  `query:"pageSize" validate:"required,numeric,min=1,max=100"`
		Desc            *bool   `query:"desc" validate:"required"`
		UserID          *string `query:"userID" validate:"omitnil,mongodb"`
		CreateStartTime *string `query:"createStartTime" validate:"omitnil,rfc3339,earlierThan=CreateEndTime"`
		CreateEndTime   *string `query:"createEndTime" validate:"omitnil,rfc3339"`
		UpdateStartTime *string `query:"updateStartTime" validate:"omitnil,rfc3339,earlierThan=UpdateEndTime"`
		UpdateEndTime   *string `query:"updateEndTime" validate:"omitnil,rfc3339"`
		Theme           *string `query:"theme" validate:""`
		Status          *string `query:"status" validate:"omitnil,instructionDataStatus"`
		Query           *string `query:"query" validate:""`
	}

	ApproveInstructionDataRequest struct {
		InstructionDataID *string `json:"instruction_data_id" validate:"required,mongodb"`
	}

	RejectInstructionDataRequest struct {
		InstructionDataID *string `json:"instruction_data_id" validate:"required,mongodb"`
		Message           *string `json:"message" validate:"required,max=1000,min=1"`
	}

	UpdateInstructionDataRequest struct {
		InstructionDataID *string `json:"instruction_data_id" validate:"required,mongodb"`
		UserID            *string `json:"user_id" validate:"omitnil,mongodb"`
		Instruction       *string `json:"instruction" validate:"omitnil,max=1000,min=1"`
		Input             *string `json:"input" validate:"omitnil,max=1000,min=1"`
		Output            *string `json:"output" validate:"omitnil,max=1000,min=1"`
		Theme             *string `json:"theme" validate:""`
		Source            *string `json:"source" validate:"omitnil,max=100"`
		Note              *string `json:"note" validate:"omitnil,max=1000"`
	}

	ExportInstructionDataRequest struct {
		Desc            *bool   `query:"desc" validate:"required"`
		UserID          *string `query:"userID" validate:"omitnil,mongodb"`
		CreateStartTime *string `query:"createStartTime" validate:"omitnil,rfc3339,earlierThan=CreateEndTime"`
		CreateEndTime   *string `query:"createEndTime" validate:"omitnil,rfc3339"`
		UpdateStartTime *string `query:"updateStartTime" validate:"omitnil,rfc3339,earlierThan=UpdateEndTime"`
		UpdateEndTime   *string `query:"updateEndTime" validate:"omitnil,rfc3339"`
		Theme           *string `query:"theme" validate:""`
		Status          *string `query:"status" validate:"omitnil,instructionDataStatus"`
	}

	DeleteInstructionDataRequest struct {
		InstructionDataID *string `query:"instructionDataID" validate:"required,mongodb"`
	}

	InsertNoticeRequest struct {
		Title      *string `json:"title" validate:"required,max=100,min=1"`
		Content    *string `json:"content" validate:"required,max=10000,min=1"`
		NoticeType *string `json:"notice_type" validate:"required,noticeType"`
	}

	UpdateNoticeRequest struct {
		NoticeID   *string `json:"notice_id" validate:"required,mongodb"`
		Title      *string `json:"title" validate:"omitnil,max=100,min=1"`
		Content    *string `json:"content" validate:"omitnil,max=10000,min=1"`
		NoticeType *string `json:"notice_type" validate:"omitnil,noticeType"`
	}

	DeleteNoticeRequest struct {
		NoticeID *string `query:"noticeID" validate:"required,mongodb"`
	}

	InsertUserRequest struct {
		Username     *string `json:"username" validate:"required,min=3,max=20"`
		Email        *string `json:"email" validate:"required,email,max=100"`
		Password     *string `json:"password" validate:"required,min=8,max=20"`
		Organization *string `json:"organization" validate:"required,max=100"`
	}

	GetUserRequest struct {
		UserID *string `query:"userID" validate:"required,mongodb"`
	}

	GetUserListRequest struct {
		Page               *int64  `query:"page" validate:"required,numeric,min=1"`
		PageSize           *int64  `query:"pageSize" validate:"required,numeric,min=1,max=100"`
		Desc               *bool   `query:"desc" validate:"required"`
		Role               *string `query:"role" validate:"omitnil,userRole"`
		LastLoginStartTime *string `query:"lastLoginStartTime" validate:"omitnil,rfc3339,earlierThan=LastLoginEndTime"`
		LastLoginEndTime   *string `query:"lastLoginEndTime" validate:"omitnil,rfc3339"`
		CreateStartTime    *string `query:"createStartTime" validate:"omitnil,rfc3339,earlierThan=CreateEndTime"`
		CreateEndTime      *string `query:"createEndTime" validate:"omitnil,rfc3339"`
		Query              *string `query:"query" validate:"omitnil,max=100"`
	}

	UpdateUserRequest struct {
		UserID       *string `json:"user_id" validate:"required"`
		Username     *string `json:"username" validate:"omitnil,min=3,max=20"`
		Email        *string `json:"email" validate:"omitnil,email"`
		Organization *string `json:"organization" validate:"omitnil,max=100"`
	}

	DeleteUserRequest struct {
		UserID *string `query:"userID" validate:"required,mongodb"`
	}

	ChangeUserPasswordRequest struct {
		UserID      *string `json:"user_id" validate:"required,mongodb"`
		NewPassword *string `json:"new_password" validate:"required,min=8,max=20"`
	}

	InsertDocumentationRequest struct {
		Title   *string `json:"title" validate:"required,max=100,min=1"`
		Content *string `json:"content" validate:"required,max=10000,min=1"`
	}

	UpdateDocumentationRequest struct {
		DocumentationID *string `json:"documentation_id" validate:"required,mongodb"`
		Title           *string `json:"title" validate:"omitnil,max=100,min=1"`
		Content         *string `json:"content" validate:"omitnil,max=10000,min=1"`
	}

	DeleteDocumentationRequest struct {
		DocumentationID *string `query:"documentationID" validate:"required,mongodb"`
	}

	GetLoginLogListRequest struct {
		Page            *int64  `query:"page" validate:"required,numeric,min=1"`
		PageSize        *int64  `query:"pageSize" validate:"required,numeric,min=1,max=100"`
		Desc            *bool   `query:"desc" validate:"required"`
		Query           *string `query:"query" validate:"omitnil,max=100"`
		CreateStartTime *string `query:"createStartTime" validate:"omitnil,rfc3339,earlierThan=CreateEndTime"`
		CreateEndTime   *string `query:"createEndTime" validate:"omitnil,rfc3339"`
	}

	GetOperationLogListRequest struct {
		Page            *int64  `query:"page" validate:"required,numeric,min=1"`
		PageSize        *int64  `query:"pageSize" validate:"required,numeric,min=1,max=100"`
		Desc            *bool   `query:"desc" validate:"required"`
		Query           *string `query:"query" validate:"omitnil,max=100"`
		Operation       *string `query:"operation" validate:"omitnil,operationType"`
		EntityType      *string `query:"entityType" validate:"omitnil,entityType"`
		Status          *string `query:"status" validate:"omitnil,operationStatus"`
		CreateStartTime *string `query:"createStartTime" validate:"omitnil,rfc3339,earlierThan=CreateEndTime"`
		CreateEndTime   *string `query:"createEndTime" validate:"omitnil,rfc3339"`
	}
)
