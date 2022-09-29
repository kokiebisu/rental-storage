import { BookingStatus } from "../enum";
import { v4 as uuid } from "uuid";
import { Amount } from "./Amount";
import { StorageItemInterface } from "../../types";

interface BookingConstructor {
  id?: string;
  status?: BookingStatus;
  amount: Amount;
  userId: string;
  listingId: string;
  createdAt?: Date;
  updatedAt?: Date;
  items: StorageItemInterface[];
}

export class Booking {
  public readonly id: string;
  public readonly status: BookingStatus;
  public readonly amount: Amount;
  public readonly userId: string;
  public readonly listingId: string;
  public readonly createdAt: Date;
  public readonly updatedAt?: Date | undefined;
  public readonly items: StorageItemInterface[];

  public constructor({
    id = uuid(),
    status = BookingStatus.CREATED,
    amount,
    userId,
    listingId,
    createdAt = new Date(),
    updatedAt,
    items,
  }: BookingConstructor) {
    this.id = id;
    this.status = status;
    this.amount = amount;
    this.userId = userId;
    this.listingId = listingId;
    this.createdAt = createdAt;
    this.updatedAt = updatedAt;
    this.items = items;
  }
}
