package entity

type Notification struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Message string `json:"message"`
	Type    string `json:"type"`
	Status  string `json:"status"`
}
