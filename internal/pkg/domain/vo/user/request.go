package user

type (
	GetDataStatisticRequest struct {
		StartDate *string `query:"start_date" validate:"datetime"`
		EndDate   *string `query:"end_date" validate:"datetime"`
	}

	GetInstructionDataRequest struct {
		InstructionDataID *string `query:"instruction_data_id" validate:"required,mongodb"`
	}

	GetInstructionDataListRequest struct {
		Page            *int64  `query:"page" validate:"required,numeric,min=1"`
		PageSize        *int64  `query:"page_size" validate:"required,numeric,min=1,max=100"`
		UpdateStartTime *string `query:"update_start_time" validate:"datetime"`
		UpdateEndTime   *string `query:"update_end_time" validate:"datetime"`
		Theme           *string `query:"theme" validate:""`
		Status          *string `query:"status" validate:""`
	}

	InsertInstructionDataRequest struct {
		Instruction *string `json:"instruction" validate:"required,max=1000"`
		Input       *string `json:"input" validate:"required,max=1000"`
		Output      *string `json:"output" validate:"required,max=1000"`
		Theme       *string `json:"theme" validate:"required"`
		Source      *string `json:"source" validate:"required,max=1000"`
		Note        *string `json:"note" validate:"required,omitempty,max=1000"`
	}

	UpdateInstructionDataRequest struct {
		InstructionDataID *string `json:"instruction_data_id" validate:"required,mongodb"`
		Instruction       *string `json:"instruction" validate:"required"`
		Input             *string `json:"input" validate:"required"`
		Output            *string `json:"output" validate:"required"`
		Theme             *string `json:"theme" validate:"required"`
		Source            *string `json:"source" validate:"required"`
		Note              *string `json:"note" validate:""`
	}

	DeleteInstructionDataRequest struct {
		InstructionDataID *string `json:"instruction_data_id" validate:"required,mongodb"`
	}
)
