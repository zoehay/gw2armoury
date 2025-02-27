interface Account {
	accountID :  string;
	accountName? :string;
	apiKey?     :string;
	password?  :  string;
	sessionID ? : string;
	session   ? :  Session;
}

interface Session {
    sessionID:   string;
    expires:  Date;
}

export default Account