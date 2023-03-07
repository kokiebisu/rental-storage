import axios, { AxiosInstance } from "axios";

export default class {
  protected client: AxiosInstance;
  public constructor() {
    this.client = axios.create({});
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
