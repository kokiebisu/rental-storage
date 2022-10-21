import { BookingStatus } from "../domain/enums";
import { Amount } from "../domain/model";
import { AmountInterface } from "./Amount";
import { ItemInterface } from "./Item";

export interface BookingConstructor {
  id?: string;
  status?: BookingStatus;
  amount: Amount;
  ownerId: number;
  listingId: number;
  createdAt?: Date;
  updatedAt?: Date;
  items: ItemInterface[];
}

export interface BookingInterface {
  id: string;
  status: string;
  amount: AmountInterface;
  ownerId: number;
  listingId: number;
  createdAt: string;
  updatedAt?: string;
  items: ItemInterface[];
}
