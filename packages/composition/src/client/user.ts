export default class {
  public baseURL: string;

  constructor() {
    const baseURL = process.env.SERVICE_API_ENDPOINT;
    if (!baseURL) {
      throw new Error("SERVICE_API_ENDPOINT not being fetched");
    }
    this.baseURL = baseURL;
  }

  public findUser(id: string) {
    const url = new URL(`${this.baseURL}/users/${id}`);
    return url.toString();
  }

  public deleteUser(id: string) {
    const url = new URL(`${this.baseURL}/users/${id}`);
    return url.toString();
  }

  // WARNING: THIS WILL NOT BE EXPOSED ON APPSYNC
  // PURELY FOR INTEGRATION TEST PURPOSES
  public createUser() {
    const url = new URL(`${this.baseURL}/users`);
    return url.toString();
  }
}
