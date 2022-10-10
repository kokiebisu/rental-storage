import { AmountInterface } from "./Amount";
import { ItemInterface } from "./Item";

export interface BookingInterface {
  id: number;
  uid: string;
  status: string;
  amount: AmountInterface;
  userId: number;
  listingId: number;
  createdAt: string;
  updatedAt?: string;
  items: ItemInterface[];
}
