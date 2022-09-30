import { BookingStatus } from "../domain/enum";
import { Amount } from "../domain/model";
import { AmountInterface } from "./Amount";
import { StorageItemInterface } from "./Item";

export interface BookingConstructor {
  id?: string;
  status?: BookingStatus;
  amount: Amount;
  guestId: string;
  listingId: string;
  createdAt?: Date;
  updatedAt?: Date;
  items: StorageItemInterface[];
}

export interface BookingInterface {
  id: string;
  status: string;
  amount: AmountInterface;
  guestId: string;
  listingId: string;
  createdAt: string;
  updatedAt?: string;
  items: StorageItemInterface[];
}
