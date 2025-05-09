package models

type Account struct {
	AccountID      string   `json:"id"`
	AccountName    *string  `json:"name"`
	GW2AccountName *string  `json:"gw2_name"`
	APIKey         *string  `json:"api_key"`
	Password       *string  `json:"password"`
	SessionID      *string  `json:"session_id"`
	Session        *Session `json:"session"`
}
