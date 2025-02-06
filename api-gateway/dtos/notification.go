package dtos

type CreateRequest struct {
	Email   string `json:"email" validate:"required,email"`
	Message string `json:"message" validate:"required"`
	Type    string `json:"type" validate:"required"`
}
