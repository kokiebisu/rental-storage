import { AmountInterface } from "./Amount";
import { StorageItemInterface } from "./Item";

export interface BookingInterface {
  id: number;
  uid: string;
  status: string;
  amount: AmountInterface;
  userId: number;
  listingId: number;
  createdAt: string;
  updatedAt?: string;
  items: StorageItemInterface[];
}
