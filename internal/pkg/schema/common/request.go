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
		NewPassword *string `json:"new_password" validate:"required, min=8, max=20"`
	}

	GetNoticeRequest struct {
		NoticeID *string `json:"notice_id" validate:"required,mongodb"`
	}

	GetNoticeListRequest struct {
		Page         *int    `json:"page" validate:"required,numeric"`
		NoticeType   *string `json:"notice_type" validate:""`
		UpdateBefore *string `json:"update_before" validate:"datetime"`
		UpdateAfter  *string `json:"update_after" validate:"datetime"`
	}

	GetDocumentationRequest struct {
		DocumentationID *string `json:"documentation_id" validate:"required"`
	}

	GetDocumentationListRequest struct {
		Page         *int    `json:"page" validate:"required,numeric"`
		UpdateBefore *string `json:"update_before" validate:"datetime"`
		UpdateAfter  *string `json:"update_after" validate:"datetime"`
	}
)
