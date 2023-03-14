export default class {
  public baseURL: string;

  constructor() {
    const baseURL = process.env.SERVICE_API_ENDPOINT;
    if (!baseURL) {
      throw new Error("SERVICE_API_ENDPOINT not being fetched");
    }
    this.baseURL = baseURL;
  }

  public getPresignedURL(filename: string) {
    const url = new URL(`${this.baseURL}/images/${filename}`);
    return url.toString();
  }
}
