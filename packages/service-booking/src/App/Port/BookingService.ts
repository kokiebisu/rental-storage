import { ItemInterface } from "../../Types";

export interface BookingService {
  makeBooking(
    amount: number,
    currency: string,
    userId: number,
    listingId: number,
    items: ItemInterface[]
  ): Promise<boolean>;
}
