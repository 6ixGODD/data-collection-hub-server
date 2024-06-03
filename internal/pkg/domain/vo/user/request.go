package user

type (
	GetDataStatisticRequest struct {
		StartDate *string `query:"startDate" validate:"omitnil,rfc3339,earlierThan=EndDate"`
		EndDate   *string `query:"endDate" validate:"omitnil,rfc3339"`
	}

	GetInstructionDataRequest struct {
		InstructionDataID *string `query:"instructionDataID" validate:"required,mongodb"`
	}

	GetInstructionDataListRequest struct {
		Page            *int64  `query:"page" validate:"required,numeric,min=1"`
		PageSize        *int64  `query:"pageSize" validate:"required,numeric,min=1,max=100"`
		UpdateStartTime *string `query:"updateStartTime" validate:"omitnil,rfc3339,earlierThan=UpdateEndTime"`
		UpdateEndTime   *string `query:"updateEndTime" validate:"omitnil,rfc3339"`
		Theme           *string `query:"theme" validate:""`
		Status          *string `query:"status" validate:"omitnil,instructionDataStatus"`
	}

	InsertInstructionDataRequest struct {
		Instruction *string `json:"instruction" validate:"required,max=1000,min=1"`
		Input       *string `json:"input" validate:"required,max=1000,min=1"`
		Output      *string `json:"output" validate:"required,max=1000,min=1"`
		Theme       *string `json:"theme" validate:""`
		Source      *string `json:"source" validate:"required,max=100"`
		Note        *string `json:"note" validate:"omitnil,max=1000"`
	}

	UpdateInstructionDataRequest struct {
		InstructionDataID *string `json:"instruction_data_id" validate:"required,mongodb"`
		Instruction       *string `json:"instruction" validate:"omitnil,max=1000,min=1"`
		Input             *string `json:"input" validate:"omitnil,max=1000,min=1"`
		Output            *string `json:"output" validate:"omitnil,max=1000,min=1"`
		Theme             *string `json:"theme" validate:""`
		Source            *string `json:"source" validate:"omitnil,max=1000"`
		Note              *string `json:"note" validate:"omitnil,max=1000"`
	}

	DeleteInstructionDataRequest struct {
		InstructionDataID *string `query:"instructionDataID" validate:"required,mongodb"`
	}
)
