package entity

import "time"

type Notification struct {
	ID        string
	Email     string
	Message   string
	Type      string
	IsSent    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
