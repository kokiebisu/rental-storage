import { BaseRestClient } from "./base";

export default class BookingRestClient extends BaseRestClient {
  public async createBooking(
    userId: string,
    spaceId: string,
    startDate: string,
    endDate: string
  ): Promise<{ uid: string }> {
    return (
      await this.client.post("/bookings", {
        userId,
        spaceId,
        startDate,
        endDate,
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
