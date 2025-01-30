interface Account {
	AccountID :  string;
	AccountName? :string;
	APIKey?     :string;
	Password?  :  string;
	SessionID ? : string;
	Session   ? :  Session;
}

interface Session {
    SessionID:   string;
    Expires:  Date;
}

export default Account