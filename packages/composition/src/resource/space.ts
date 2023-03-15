export default class {
  public baseURL: string;

  constructor() {
    const baseURL = process.env.SERVICE_API_ENDPOINT;
    if (!baseURL) {
      throw new Error("SERVICE_API_ENDPOINT not being fetched");
    }
    this.baseURL = baseURL;
  }

  public createSpace() {
    const url = new URL(`${this.baseURL}/spaces`).toString();
    return url.toString();
  }

  public findSpace(id: string) {
    const url = new URL(`${this.baseURL}/spaces/${id}`);
    return url.toString();
  }

  public deleteSpace(id: string) {
    const url = new URL(`${this.baseURL}/spaces/${id}`);
    return url.toString();
  }

  public findSpaces(param: {
    userId?: string;
    limit?: number;
    offset?: number;
  }) {
    let url = new URL(`${this.baseURL}/spaces`);
    url = this.attachQueryParams(url, param);
    return url.toString();
  }

  private attachQueryParams(
    url: URL,
    {
      userId,
      limit,
      offset,
    }: { userId?: string; limit?: number; offset?: number }
  ) {
    if (userId) {
      url.searchParams.set("userId", userId);
    }
    if (offset !== undefined) {
      url.searchParams.set("offset", offset.toString());
    }
    if (limit !== undefined) {
      url.searchParams.set("limit", limit.toString());
    }
    return url;
  }
}
