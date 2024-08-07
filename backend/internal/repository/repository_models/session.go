package repositorymodels

import "time"

type Session struct {
	SessionID string
	Expires   time.Time
}

func (session Session) isExpired() bool {
	return session.Expires.Before(time.Now())
}
