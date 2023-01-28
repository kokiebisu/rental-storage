import { BaseRestClient } from "./base";

export default class BookingRestClient extends BaseRestClient {
  public async findAllCreatedBookingsByUserId(userId: string) {
    const response = await this.client.get(`/bookings/${userId}`);
    return response.data;
  }

  public async makeBooking(
    amount: number,
    currency: string,
    userId: string,
    listingId: string,
    items: unknown
  ) {
    const response = await this.client.post("/bookings", {
      amount,
      currency,
      userId,
      listingId,
      items,
    });
    return response.data;
  }
}
