import { BaseRestClient } from "./base";

export default class BookingRestClient extends BaseRestClient {
  public async findAllCreatedBookingsByUserId(
    userId: string
  ): Promise<{ bookings: Booking[] } | undefined> {
    const response = await this.client.get(`/bookings/${userId}`);
    return response.data;
  }

  public async makeBooking(
    userId: string,
    listingId: string,
    items: BookingItem[]
  ): Promise<{ uid: string }> {
    const response = await this.client.post("/bookings", {
      userId,
      listingId,
      items,
    });
    return response.data;
  }

  public async findBookingById(
    bookingId: string
  ): Promise<{ booking: Booking } | undefined> {
    try {
      const response = await this.client.get(`/bookings/${bookingId}`);
      return response.data;
    } catch (err) {
      console.error(err);
    }
  }
}
