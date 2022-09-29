import { Currency, StorageItemInterface } from "@rental-storage-project/common";
import { Booking } from "../../domain/model";

export interface BookingService {
  makeBooking(
    amount: number,
    currency: Currency,
    userId: string,
    listingId: string,
    items: StorageItemInterface[]
  ): Promise<boolean>;
}
