import { Account, APIAccount, APIAccountToAccount } from "../models/Account";
import { BagItem, APIBagItem, APIBagItemToBagItem } from "../models/BagItem";

export interface ClientInterface {
  getBagItems(): BagItem[];
}

export class Client {
  baseURL: string;

  constructor(mode: string) {
    if (mode === "development") {
      this.baseURL = "http://localhost:8000";
    } else {
      this.baseURL = "";
    }
  }

  async clientGet(endpoint: string): Promise<any> {
    try {
      let response = await fetch(endpoint, {
        credentials: "include",
      });
      if (response.ok) {
        let responseJSON = await response.json();
        return responseJSON;
      }
    } catch (error) {
      console.log(error);
    }
  }

  async clientPost(endpoint: string, body: any): Promise<any> {
    try {
      let response = await fetch(endpoint, {
        credentials: "include",
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: body,
      });

      if (response.ok) {
        let responseJSON = await response.json();
        return responseJSON;
      }
    } catch (error) {
      console.log(error);
    }
  }

  async getBagItems(): Promise<BagItem[]> {
    let endpoint: string = `${this.baseURL}/account/inventory`;
    let response: unknown = await this.clientGet(endpoint);
    let apiBagItems = response as APIBagItem[];
    let bagItems: BagItem[] = apiBagItems.map(APIBagItemToBagItem);
    if (bagItems) {
      return bagItems;
    } else {
      return [];
    }
  }

  async postAPIKey(key: string): Promise<Account | null> {
    let body = JSON.stringify({
      APIKey: key,
    });

    let endpoint: string = `${this.baseURL}/apikeys`;
    let response: Response = await this.clientPost(endpoint, body);
    if (response.data) {
      return response.data;
    } else {
      return null;
    }
  }

  async getAccount(): Promise<Account | null> {
    let endpoint: string = `${this.baseURL}/account/info`;
    let response: unknown = await this.clientGet(endpoint);
    let apiAccount = response as APIAccount;
    let account: Account = APIAccountToAccount(apiAccount);
    if (account) {
      return account;
    } else {
      return null;
    }
  }
}

interface Response {
  data: any;
}
