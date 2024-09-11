package dbmodels

type DBAccount struct {
	AccountID   string `gorm:"primaryKey"` // The unique persistent GW2 API account GUID
	AccountName *string
	APIKey      *string
	Password    *string
	SessionID   *string    `gorm:"index"`
	Session     *DBSession `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
