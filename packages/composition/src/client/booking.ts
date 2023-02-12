import { BaseRestClient } from "./base";

export default class BookingRestClient extends BaseRestClient {
  public async createBooking(
    userId: string,
    spaceId: string,
    items: BookingItem[]
  ): Promise<{ uid: string }> {
    return (
      await this.client.post("/bookings", {
        userId,
        spaceId,
        items,
      })
    ).data;
  }

  public async findBooking(id: string) {
    return (await this.client.get(`/bookings/${id}`)).data;
  }

  public async findBookings(spaceId: string) {
    return (await this.client.get(`/bookings?spaceId=${spaceId}`)).data;
  }
}
