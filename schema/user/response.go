package user

type GetDataStatisticResponse struct {
	Total              int64                `json:"total"`
	PendingCount       int64                `json:"pending_count"`
	ApprovedCount      int64                `json:"approved_count"`
	RejectedCount      int64                `json:"rejected_count"`
	ThemeCount         map[string]int64     `json:"theme_count"`
	TimeRangeStatistic []timeRangeStatistic `json:"time_range_statistic"`
}

type timeRangeStatistic struct {
	Date          string           `json:"date"`
	Total         int64            `json:"total"`
	PendingCount  int64            `json:"pending_count"`
	ApprovedCount int64            `json:"approved_count"`
	RejectedCount int64            `json:"rejected_count"`
	ThemeCount    map[string]int64 `json:"theme_count"`
}

type GetInstructionDataResponse struct {
	RowUUID   string            `json:"row_uuid"`
	Row       instructionRow    `json:"row"`
	Theme     string            `json:"theme"`
	Source    string            `json:"source"`
	Note      string            `json:"note"`
	Status    instructionStatus `json:"status"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
}

type instructionRow struct {
	Instruction string `json:"instruction"`
	Input       string `json:"input"`
	Output      string `json:"output"`
}

type instructionStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type GetInstructionListResponse struct {
	Total               int64                        `json:"total"`
	InstructionDataList []GetInstructionDataResponse `json:"instruction_data_list"`
}
