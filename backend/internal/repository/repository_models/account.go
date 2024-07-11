package repositorymodels

type Account struct {
	AccountID   string // The unique persistent GW2 API account GUID
	AccountName *string
	APIKey      *string
	Password    *string
	Session     *Session
}
