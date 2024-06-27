package handlers

import (
	"github.com/zoehay/gw2armoury/backend/internal/repository"
)

type SessionHandler struct {
	SessionRepository repository.GORMSessionRepository
}

func NewSessionHandler(GORMSessionRepository repository.GORMSessionRepository) *SessionHandler {
	return &SessionHandler{
		SessionRepository: GORMSessionRepository,
	}
}
