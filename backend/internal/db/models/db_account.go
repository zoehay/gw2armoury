package dbmodels

import (
	"time"

	"github.com/zoehay/gw2armoury/backend/internal/api/models"
)

type DBAccount struct {
	AccountID      string `gorm:"primaryKey"` // The unique persistent GW2 API account GUID
	LastCrawl      *time.Time
	AccountName    *string
	GW2AccountName *string
	APIKey         *string
	Password       *string
	SessionID      *string    `gorm:"index"`
	Session        *DBSession `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (dbAccount DBAccount) ToAccount() models.Account {
	return models.Account{
		AccountID:      dbAccount.AccountID,
		AccountName:    dbAccount.AccountName,
		GW2AccountName: dbAccount.GW2AccountName,
		APIKey:         dbAccount.APIKey,
		Password:       dbAccount.Password,
		SessionID:      dbAccount.SessionID,
		Session:        (*models.Session)(dbAccount.Session),
	}
}
