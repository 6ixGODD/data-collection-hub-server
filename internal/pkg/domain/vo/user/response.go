package user

type (
	GetDataStatisticResponse struct {
		Total              int64                 `json:"total"`
		PendingCount       int64                 `json:"pending_count"`
		ApprovedCount      int64                 `json:"approved_count"`
		RejectedCount      int64                 `json:"rejected_count"`
		ThemeCount         map[string]int64      `json:"theme_count"`
		TimeRangeStatistic []*TimeRangeStatistic `json:"time_range_statistic"`
	}

	TimeRangeStatistic struct {
		Date          string           `json:"date"`
		Total         int64            `json:"total"`
		PendingCount  int64            `json:"pending_count"`
		ApprovedCount int64            `json:"approved_count"`
		RejectedCount int64            `json:"rejected_count"`
		ThemeCount    map[string]int64 `json:"theme_count"`
	}

	GetInstructionDataResponse struct {
		InstructionDataID string `json:"instruction_data_id"`
		Row               struct {
			Instruction string `json:"instruction"`
			Input       string `json:"input"`
			Output      string `json:"output"`
		} `json:"row"`
		Theme  string `json:"theme"`
		Source string `json:"source"`
		Note   string `json:"note"`
		Status struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"status"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	GetInstructionDataListResponse struct {
		Total               int64                         `json:"total"`
		InstructionDataList []*GetInstructionDataResponse `json:"instruction_data_list"`
	}
)
