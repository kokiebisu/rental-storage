import { BaseRestClient } from "./base";

export default class BookingRestClient extends BaseRestClient {
  // not used yet
  // public async findAllCreatedBookingsByUserId(
  //   userId: string
  // ): Promise<{ bookings: Booking[] } | undefined> {
  //   const response = await this.client.get(`/bookings/${userId}`);
  //   return response.data;
  // }

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
}
