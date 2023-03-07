export default class {
  public static createBooking() {
    return "/bookings";
  }

  public static findBooking(id: string) {
    return `/bookings/${id}`;
  }

  public static deleteBooking(id: string) {
    return `/bookings/${id}`;
  }

  public static findAllBookings(param: { spaceId?: string; userId?: string }) {
    const queryParam = this.buildQueryParam(param);
    return `/bookings?${queryParam}`;
  }

  public static findPendingBookings(param: {
    spaceId?: string;
    userId?: string;
  }) {
    const queryParam = this.buildQueryParam(param);
    return `/bookings?${queryParam}&status=pending`;
  }

  public static findApprovedBookings(param: {
    spaceId?: string;
    userId?: string;
  }) {
    const queryParam = this.buildQueryParam(param);
    return `/bookings?${queryParam}&status=approved`;
  }

  private static buildQueryParam(params: { [key: string]: string }) {
    return Object.keys(params)
      .map((key) => `${key}=${params[key]}`)
      .join("&");
  }
}
