package admin

type (
	GetDataStatisticResponse struct {
		Total              int                  `json:"total"`
		PendingCount       int                  `json:"pending_count"`
		ApprovedCount      int                  `json:"approved_count"`
		RejectedCount      int                  `json:"rejected_count"`
		ThemeCount         map[string]int       `json:"theme_count"`
		TimeRangeStatistic []timeRangeStatistic `json:"time_range_statistic"`
	}

	timeRangeStatistic struct {
		Date          string         `json:"date"`
		Total         int            `json:"total"`
		PendingCount  int            `json:"pending_count"`
		ApprovedCount int            `json:"approved_count"`
		RejectedCount int            `json:"rejected_count"`
		ThemeCount    map[string]int `json:"theme_count"`
	}

	GetUserStatisticResponse struct {
		Username string        `json:"username"`
		Data     userStatistic `json:"data"`
	}

	userStatistic struct {
		Total         int `json:"total"`
		PendingCount  int `json:"pending_count"`
		ApprovedCount int `json:"approved_count"`
		RejectedCount int `json:"rejected_count"`
	}

	GetUserStatisticListResponse struct {
		Total             int                        `json:"total"`
		UserStatisticList []GetUserStatisticResponse `json:"user_statistic_list"`
	}

	GetInstructionDataResponse struct {
		InstructionDataID string `json:"instruction_data_id"`
		UserID            string `json:"user_id"`
		Username          string `json:"username"`
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
		Total               int                          `json:"total"`
		InstructionDataList []GetInstructionDataResponse `json:"instruction_data_list"`
	}

	GetUserResponse struct {
		UserID       string `json:"_id"`
		UUID         string `json:"uuid"`
		Username     string `json:"username"`
		Email        string `json:"email"`
		Password     string `json:"password"`
		Role         string `json:"role"`
		Organization string `json:"organization"`
		LastLogin    string `json:"last_login"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
	}

	GetUserListResponse struct {
		Total    int               `json:"total"`
		UserList []GetUserResponse `json:"user_list"`
	}

	GetLoginLogResponse struct {
		LoginLogID string `json:"_id"`
		UserID     string `json:"user_id"`
		Username   string `json:"username"`
		Email      string `json:"email"`
		IPAddress  string `json:"ip_address"`
		UserAgent  string `json:"user_agent"`
		CreatedAt  string `json:"created_at"`
	}

	GetLoginLogListResponse struct {
		Total        int                   `json:"total"`
		LoginLogList []GetLoginLogResponse `json:"login_log_list"`
	}

	GetOperationLogResponse struct {
		OperationLogID string `json:"_id"`
		UserID         string `json:"user_id"`
		Username       string `json:"username"`
		Email          string `json:"email"`
		IPAddress      string `json:"ip_address"`
		UserAgent      string `json:"user_agent"`
		Operation      string `json:"operation"`
		EntityUUID     string `json:"entity_uuid"`
		CreatedAt      string `json:"created_at"`
	}

	GetOperationLogListResponse struct {
		Total            int                       `json:"total"`
		OperationLogList []GetOperationLogResponse `json:"operation_log_list"`
	}

	GetErrorLogResponse struct {
		ErrorLogID     string `json:"_id"`
		UserID         string `json:"user_id"`
		Username       string `json:"username"`
		IPAddress      string `json:"ip_address"`
		UserAgent      string `json:"user_agent"`
		RequestURI     string `json:"request_uri"`
		RequestMethod  string `json:"request_method"`
		RequestPayload string `json:"request_payload"`
		ErrorCode      string `json:"error_code"`
		ErrorMsg       string `json:"error_msg"`
		Stack          string `json:"stack"`
		CreatedAt      string `json:"created_at"`
	}

	GetErrorLogListResponse struct {
		Total        int                   `json:"total"`
		ErrorLogList []GetErrorLogResponse `json:"error_log_list"`
	}
)
