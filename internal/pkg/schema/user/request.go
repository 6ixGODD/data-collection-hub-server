package user

type GetDataStatisticRequest struct {
	StartDate string `json:"start_date" validate:""`
	EndDate   string `json:"end_date" validate:""`
}

type GetInstructionDataRequest struct {
	InstructionDataID string `json:"instruction_data_id" validate:"required"`
}

type GetInstructionListRequest struct {
	Page         int    `json:"page" validate:"required"`
	UpdateBefore string `json:"update_before" validate:""`
	UpdateAfter  string `json:"update_after" validate:""`
	Theme        string `json:"theme" validate:""`
	Status       int    `json:"status" validate:""`
}

type InsertInstructionRequest struct {
	Instruction string `json:"instruction" validate:"required"`
	Input       string `json:"input" validate:"required"`
	Output      string `json:"output" validate:"required"`
	Theme       string `json:"theme" validate:"required"`
	Source      string `json:"source" validate:"required"`
	Note        string `json:"note" validate:""`
}

type UpdateInstructionRequest struct {
	InstructionDataID string `json:"instruction_data_id" validate:"required"`
	Instruction       string `json:"instruction" validate:"required"`
	Input             string `json:"input" validate:"required"`
	Output            string `json:"output" validate:"required"`
	Theme             string `json:"theme" validate:"required"`
	Source            string `json:"source" validate:"required"`
	Note              string `json:"note" validate:""`
}

type DeleteInstructionRequest struct {
	InstructionDataID string `json:"instruction_data_id" validate:"required"`
}
