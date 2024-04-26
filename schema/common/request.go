package common

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

type GetNoticeRequest struct {
	NoticeID string `json:"notice_id" validate:"required"`
}

type GetNoticeListRequest struct {
	Page         int    `json:"page" validate:"required"`
	Type         string `json:"type" validate:""`
	UpdateBefore string `json:"update_before" validate:""`
	UpdateAfter  string `json:"update_after" validate:""`
}

type GetDocumentationRequest struct {
	DocumentationID string `json:"documentation_id" validate:"required"`
}

type GetDocumentationListRequest struct {
	Page         int    `json:"page" validate:"required"`
	UpdateBefore string `json:"update_before" validate:""`
	UpdateAfter  string `json:"update_after" validate:""`
}
