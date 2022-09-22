import { BookingStatus } from "../enum";
import { Amount } from "@rental-storage-project/common";

export class Booking {
  private _status: BookingStatus;
  private _amount: Amount;
  private _userId: string;
  private _listingId: string;
  private _createdAt: Date;
  private _updatedAt?: Date;

  public constructor(
    status: BookingStatus,
    amount: Amount,
    userId: string,
    listingId: string
  ) {
    this._status = status;
    this._amount = amount;
    this._userId = userId;
    this._listingId = listingId;
    this._createdAt = new Date();
  }
}
