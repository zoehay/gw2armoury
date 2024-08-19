package repositorymodels

type DBAccount struct {
	AccountID   string // The unique persistent GW2 API account GUID
	AccountName *string
	APIKey      *string
	Password    *string
	Session     *DBSession
}
