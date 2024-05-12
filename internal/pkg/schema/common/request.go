package common

type (
	LoginRequest struct {
		Email    *string `json:"email" validate:"required,email"`
		Password *string `json:"password" validate:"required"`
	}

	RefreshTokenRequest struct {
		RefreshToken *string `json:"refresh_token" validate:"required"`
	}

	ChangePasswordRequest struct {
		OldPassword *string `json:"old_password" validate:"required"`
		NewPassword *string `json:"new_password" validate:"required,min=8,max=20"`
	}

	GetNoticeRequest struct {
		NoticeID *string `json:"notice_id" validate:"required,mongodb"`
	}

	GetNoticeListRequest struct {
		Page            *int64  `json:"page" validate:"required,numeric,min=1"`
		PageSize        *int64  `json:"page_size" validate:"required,numeric,min=1,max=100"`
		NoticeType      *string `json:"notice_type" validate:""`
		UpdateStartTime *string `json:"update_start_time" validate:"datetime"`
		UpdateEndTime   *string `json:"update_end_time" validate:"datetime"`
	}

	GetDocumentationRequest struct {
		DocumentationID *string `json:"documentation_id" validate:"required"`
	}

	GetDocumentationListRequest struct {
		Page            *int64  `json:"page" validate:"required,numeric,min=1"`
		PageSize        *int64  `json:"page_size" validate:"required,numeric,min=1,max=100"`
		UpdateStartTime *string `json:"update_start_time" validate:"datetime"`
		UpdateEndTime   *string `json:"update_end_time" validate:"datetime"`
	}
)
