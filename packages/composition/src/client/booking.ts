import { URL } from "url";

export default class {
  public baseURL: string;

  constructor() {
    const baseURL = process.env.SERVICE_API_ENDPOINT;
    if (!baseURL) {
      throw new Error("SERVICE_API_ENDPOINT not being fetched");
    }
    this.baseURL = baseURL;
  }

  public createBooking() {
    const url = new URL(`${this.baseURL}/bookings`).toString();
    return url.toString();
  }

  public findBooking(id: string) {
    const url = new URL(`${this.baseURL}/bookings/${id}`);
    return url.toString();
  }

  public deleteBooking(id: string) {
    const url = new URL(`${this.baseURL}/bookings/${id}`);
    return url.toString();
  }

  public findAllBookings(param: { spaceId?: string; userId?: string }) {
    let url = new URL(`${this.baseURL}/bookings`);
    url = this.attachQueryParam(url, { ...param });
    return url.toString();
  }

  public findPendingBookings(param: { spaceId?: string; userId?: string }) {
    let url = new URL(`${this.baseURL}/bookings`);
    url = this.attachQueryParam(url, { ...param, approved: false });
    return url.toString();
  }

  public findApprovedBookings(param: { spaceId?: string; userId?: string }) {
    let url = new URL(`${this.baseURL}/bookings`);
    url = this.attachQueryParam(url, { ...param, approved: true });
    return url.toString();
  }

  private attachQueryParam(
    url: URL,
    param: { spaceId?: string; userId?: string; approved?: boolean }
  ) {
    if (param.spaceId) {
      url.searchParams.set("spaceId", param.spaceId);
    }
    if (param.userId) {
      url.searchParams.set("userId", param.userId);
    }
    if (param.approved) {
      url.searchParams.set("status", "approved");
    } else if (param.approved === false) {
      url.searchParams.set("status", "pending");
    }
    return url;
  }
}
