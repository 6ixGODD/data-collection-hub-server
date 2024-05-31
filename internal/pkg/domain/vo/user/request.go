package user

type (
	GetDataStatisticRequest struct {
		StartDate *string `query:"start_date" validate:"omitnil,rfc3339,earlierThan=EndDate"`
		EndDate   *string `query:"end_date" validate:"omitnil,rfc3339"`
	}

	GetInstructionDataRequest struct {
		InstructionDataID *string `query:"instruction_data_id" validate:"required,mongodb"`
	}

	GetInstructionDataListRequest struct {
		Page            *int64  `query:"page" validate:"required,numeric,min=1"`
		PageSize        *int64  `query:"page_size" validate:"required,numeric,min=1,max=100"`
		UpdateStartTime *string `query:"update_start_time" validate:"omitnil,rfc3339,earlierThan=UpdateEndTime"`
		UpdateEndTime   *string `query:"update_end_time" validate:"omitnil,rfc3339"`
		Theme           *string `query:"theme" validate:"omitnil,InstructionDataTheme"`
		Status          *string `query:"status" validate:"omitnil,instructionDataStatus"`
	}

	InsertInstructionDataRequest struct {
		Instruction *string `json:"instruction" validate:"required,max=1000,min=1"`
		Input       *string `json:"input" validate:"required,max=1000,min=1"`
		Output      *string `json:"output" validate:"required,max=1000,min=1"`
		Theme       *string `json:"theme" validate:"required,InstructionDataTheme"`
		Source      *string `json:"source" validate:"required,max=100"`
		Note        *string `json:"note" validate:"omitnil,max=1000"`
	}

	UpdateInstructionDataRequest struct {
		InstructionDataID *string `json:"instruction_data_id" validate:"required,mongodb"`
		Instruction       *string `json:"instruction" validate:"required,max=1000,min=1"`
		Input             *string `json:"input" validate:"required,max=1000,min=1"`
		Output            *string `json:"output" validate:"required,max=1000,min=1"`
		Theme             *string `json:"theme" validate:"required,InstructionDataTheme"`
		Source            *string `json:"source" validate:"required,max=1000"`
		Note              *string `json:"note" validate:"omitnil,max=1000"`
	}

	DeleteInstructionDataRequest struct {
		InstructionDataID *string `json:"instruction_data_id" validate:"required,mongodb"`
	}
)
