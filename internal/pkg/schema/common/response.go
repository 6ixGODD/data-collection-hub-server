package common

type (
	meta struct {
		UserID   string `json:"user_id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Role     string `json:"role"`
	}

	LoginResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int64  `json:"expires_in"`
		Meta         meta   `json:"meta"`
	}

	RefreshTokenResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int64  `json:"expires_in"`
		Meta         meta   `json:"meta"`
	}

	GetNoticeResponse struct {
		NoticeID  string `json:"_id"`
		Title     string `json:"title"`
		Content   string `json:"content"`
		Type      string `json:"type"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	noticeSummary struct {
		NoticeID  string `json:"_id"`
		Title     string `json:"title"`
		Type      string `json:"type"`
		CreatedAt string `json:"created_at"`
	}

	GetNoticeListResponse struct {
		Total             int64           `json:"total"`
		NoticeSummaryList []noticeSummary `json:"notice_summary_list"`
	}

	GetDocumentationResponse struct {
		DocumentID string `json:"_id"`
		Title      string `json:"title"`
		Content    string `json:"content"`
		CreatedAt  string `json:"created_at"`
		UpdatedAt  string `json:"updated_at"`
	}

	documentationSummary struct {
		DocumentID string `json:"_id"`
		Title      string `json:"title"`
		CreatedAt  string `json:"created_at"`
	}

	GetDocumentationListResponse struct {
		Total                    int                    `json:"total"`
		DocumentationSummaryList []documentationSummary `json:"documentation_summary_list"`
	}

	GetProfileResponse struct {
		UserID       string `json:"user_id"`
		Username     string `json:"username"`
		Email        string `json:"email"`
		Role         string `json:"role"`
		Organization string `json:"organization"`
		LastLogin    string `json:"last_login"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
	}

	ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)
