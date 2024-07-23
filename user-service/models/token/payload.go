package model

import "time"

type Payload struct {
	UserID    string `json:"user_id"`
	Type      string `json:"type"`
	Scope     string `json:"scope,omitempty"`
	CreatedAt time.Time
	ExpiredAt time.Time
}
