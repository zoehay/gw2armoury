import BagItem from "../models/BagItem";

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

      console.log(response)
      if (response.ok) {
        let responseJSON = await response.json();
        return responseJSON;
      }
    } catch (error) {
      console.log(error);
    }
  }
  async getBagItems(): Promise<BagItem[]> {
    let endpoint: string = `${this.baseURL}/items`;
    let response: Response = await this.clientGet(endpoint);
    console.log(response)
    if (response.data) {
      return response.data;
    } else {
      return [];
    }
  }
}

interface Response {
  data: any;
}
