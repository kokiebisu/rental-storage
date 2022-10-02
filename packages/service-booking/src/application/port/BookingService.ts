import { StorageItemInterface } from "../../types";

export interface BookingService {
  makeBooking(
    amount: number,
    currency: string,
    userId: number,
    listingId: number,
    items: StorageItemInterface[]
  ): Promise<boolean>;
}
