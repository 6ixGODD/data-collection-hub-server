package common

type (
	LoginRequest struct {
		Email    *string `json:"email" validate:"required,email"`
		Password *string `json:"password" validate:"required"`
	}

	RefreshTokenRequest struct {
		RefreshToken *string `json:"refresh_token" validate:"required,jwt"`
	}

	ChangePasswordRequest struct {
		OldPassword *string `json:"old_password" validate:"required"`
		NewPassword *string `json:"new_password" validate:"required,min=8,max=20"`
	}

	GetNoticeRequest struct {
		NoticeID *string `query:"notice_id" validate:"required,mongodb"`
	}

	GetNoticeListRequest struct {
		Page            *int64  `query:"page" validate:"required,numeric,min=1"`
		PageSize        *int64  `query:"page_size" validate:"required,numeric,min=1,max=100"`
		NoticeType      *string `query:"notice_type" validate:"omitnil,noticeType"`
		UpdateStartTime *string `query:"update_start_time" validate:"omitnil,rfc3339,earlierThan=UpdateEndTime"`
		UpdateEndTime   *string `query:"update_end_time" validate:"omitnil,rfc3339"`
	}

	GetDocumentationRequest struct {
		DocumentationID *string `json:"documentation_id" validate:"required"`
	}

	GetDocumentationListRequest struct {
		Page            *int64  `query:"page" validate:"required,numeric,min=1"`
		PageSize        *int64  `query:"page_size" validate:"required,numeric,min=1,max=100"`
		UpdateStartTime *string `query:"update_start_time" validate:"omitnil,rfc3339,earlierThan=UpdateEndTime"`
		UpdateEndTime   *string `query:"update_end_time" validate:"omitnil,rfc3339"`
	}
)
