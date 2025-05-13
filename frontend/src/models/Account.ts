export interface Account {
  accountID: string;
  accountName?: string;
  apiKey?: string;
  password?: string;
  sessionID?: string;
  session?: Session;
}

export interface APIAccount {
  id: string;
  account_name?: string;
  name?: string;
  gw2_name?: string;
  api_key?: string;
  password?: string;
  session_id?: string;
}

export function APIAccountToAccount(apiAccount: APIAccount): Account {
  return {
    accountID: apiAccount.id,
    accountName: apiAccount.account_name,
    apiKey: apiAccount.api_key,
    password: apiAccount.password,
    sessionID: apiAccount.session_id,
  };
}

interface Session {
  sessionID: string;
  expires: Date;
}

export default Account;
