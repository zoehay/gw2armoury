package dbmodels

import "time"

type DBSession struct {
	SessionID string
	Expires   time.Time
}

// func (session Session) isExpired() bool {
// 	return session.Expires.Before(time.Now())
// }
