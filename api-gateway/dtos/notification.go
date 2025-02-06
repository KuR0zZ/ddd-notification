package dtos

type CreateRequest struct {
	Email   string `json:"email"`
	Message string `json:"message"`
	Type    string `json:"type"`
}
