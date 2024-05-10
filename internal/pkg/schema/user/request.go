package user

type (
	GetDataStatisticRequest struct {
		StartDate *string `json:"start_date" validate:"datetime"`
		EndDate   *string `json:"end_date" validate:"datetime"`
	}

	GetInstructionDataRequest struct {
		InstructionDataID *string `json:"instruction_data_id" validate:"required,mongodb"`
	}

	GetInstructionDataListRequest struct {
		Page         *int    `json:"page" validate:"required,numeric"`
		UpdateBefore *string `json:"update_before" validate:"datetime"`
		UpdateAfter  *string `json:"update_after" validate:"datetime"`
		Theme        *string `json:"theme" validate:""`
		Status       *string `json:"status" validate:""`
	}

	InsertInstructionDataRequest struct {
		Instruction *string `json:"instruction" validate:"required"`
		Input       *string `json:"input" validate:"required"`
		Output      *string `json:"output" validate:"required"`
		Theme       *string `json:"theme" validate:"required"`
		Source      *string `json:"source" validate:"required"`
		Note        *string `json:"note" validate:""`
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
