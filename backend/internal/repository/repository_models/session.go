package repositorymodels

import "time"

type Session struct {
	SessionID string
	Expires   time.Time
}
