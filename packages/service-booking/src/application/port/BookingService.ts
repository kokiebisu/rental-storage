import { Currency } from "../../domain/enum";
import { StorageItemInterface } from "../../types";

export interface BookingService {
  makeBooking(
    amount: number,
    currency: Currency,
    userId: string,
    listingId: string,
    items: StorageItemInterface[]
  ): Promise<boolean>;
}
