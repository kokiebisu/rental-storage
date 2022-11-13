import axios, { AxiosInstance } from "axios";

export class BaseRestClient {
  protected client: AxiosInstance;
  public constructor() {
    const baseURL = process.env.SERVICE_API_ENDPOINT;
    if (!baseURL) {
      throw new Error("SERVICE_API_ENDPOINT not being fetched");
    }
    this.client = axios.create({
      baseURL,
    });
  }
}
