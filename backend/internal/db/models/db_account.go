package dbmodels

type DBAccount struct {
	AccountID   string `gorm:"primaryKey"` // The unique persistent GW2 API account GUID
	AccountName *string
	APIKey      *string
	Password    *string
	Session     *DBSession `gorm:"foreignKey:SessionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
