import { AmountInterface } from "./Amount";
import { ItemInterface } from "./Item";

export interface BookingInterface {
  id: number;
  status: string;
  amount: AmountInterface;
  userId: number;
  listingId: string;
  createdAt: string;
  updatedAt?: string;
  items: ItemInterface[];
}
