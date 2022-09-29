import { AmountInterface } from "./Amount";
import { StorageItemInterface } from "./Item";

export interface BookingInterface {
  id: string;
  status: string;
  amount: AmountInterface;
  userId: string;
  listingId: string;
  createdAt: string;
  updatedAt?: string;
  items: StorageItemInterface[];
}
