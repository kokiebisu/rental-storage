import { BookingItem } from "../adapter/usecase/makeBooking";
import { BaseRestClient } from "./base";

export default class BookingRestClient extends BaseRestClient {
  public async findAllCreatedBookingsByUserId(userId: string) {
    const response = await this.client.get(`/bookings/${userId}`);
    return response.data;
  }

  public async makeBooking(
    userId: string,
    listingId: string,
    items: BookingItem[]
  ) {
    try {
      const response = await this.client.post("/bookings", {
        userId,
        listingId,
        items,
      });
      return response.data;
    } catch (err) {
      console.error(err);
    }
  }

  public async findBookingById(bookingId: string) {
    try {
      const response = await this.client.get(`/bookings/${bookingId}`);
      return response.data;
    } catch (err) {
      console.error(err);
    }
  }
}
