import { BookingStatus } from "../enum";
import { v4 as uuid } from "uuid";
import { Amount } from "./Amount";
import { BookingConstructor, ItemInterface } from "../types";

export class Booking {
  public readonly id: string;
  public readonly status: BookingStatus;
  public readonly amount: Amount;
  public readonly ownerId: number;
  public readonly listingId: number;
  public readonly createdAt: Date;
  public readonly updatedAt?: Date | undefined;
  public readonly items: ItemInterface[];

  public constructor({
    id = uuid(),
    status = BookingStatus.CREATED,
    amount,
    ownerId,
    listingId,
    createdAt = new Date(),
    updatedAt,
    items,
  }: BookingConstructor) {
    this.id = id;
    this.status = status;
    this.amount = amount;
    this.ownerId = ownerId;
    this.listingId = listingId;
    this.createdAt = createdAt;
    this.updatedAt = updatedAt;
    this.items = items;
  }
}
