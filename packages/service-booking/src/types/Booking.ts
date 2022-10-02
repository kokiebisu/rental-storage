import { BookingStatus } from "../domain/enum";
import { Amount } from "../domain/model";
import { AmountInterface } from "./Amount";
import { StorageItemInterface } from "./Item";

export interface BookingConstructor {
  id?: string;
  status?: BookingStatus;
  amount: Amount;
  guestId: number;
  listingId: number;
  createdAt?: Date;
  updatedAt?: Date;
  items: StorageItemInterface[];
}

export interface BookingInterface {
  id: string;
  status: string;
  amount: AmountInterface;
  guestId: number;
  listingId: number;
  createdAt: string;
  updatedAt?: string;
  items: StorageItemInterface[];
}
