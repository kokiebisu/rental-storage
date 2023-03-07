import axios, { AxiosInstance } from "axios";

export default class {
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

  public async get<T>(path: string) {
    return this.client.get<T>(path);
  }

  public async post<T, D>(path: string, data: D) {
    return this.client.post<T>(path, data);
  }

  public async delete<T>(path: string) {
    return this.client.delete<T>(path);
  }

  public async put<T, D>(path: string, data: D) {
    return this.client.put<T>(path, data);
  }
}
