package models

import "time"

type Session struct {
	SessionID string    `json:"session_id"`
	Expires   time.Time `json:"expires"`
}
