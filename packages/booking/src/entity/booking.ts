import { BookingStatus } from "../enum";
import { Amount } from "@rental-storage-project/common";
import { v4 as uuid } from "uuid";

interface BookingConstructor {
  id?: string;
  status?: BookingStatus;
  amount: Amount;
  userId: string;
  listingId: string;
  createdAt?: Date;
  updatedAt?: Date;
}

export class Booking {
  private _id: string;

  private _status: BookingStatus;

  private _amount: Amount;

  private _userId: string;

  private _listingId: string;

  private _createdAt: Date;

  private _updatedAt?: Date | undefined;

  public constructor({
    id = uuid(),
    status = BookingStatus.CREATED,
    amount,
    userId,
    listingId,
    createdAt = new Date(),
    updatedAt,
  }: BookingConstructor) {
    this._id = id;
    this._status = status;
    this._amount = amount;
    this._userId = userId;
    this._listingId = listingId;
    this._createdAt = createdAt;
    this._updatedAt = updatedAt;
  }

  public get id(): string {
    return this._id;
  }
  public set id(value: string) {
    this._id = value;
  }

  public get status(): BookingStatus {
    return this._status;
  }
  public set status(value: BookingStatus) {
    this._status = value;
  }

  public get amount(): Amount {
    return this._amount;
  }
  public set amount(value: Amount) {
    this._amount = value;
  }

  public get userId(): string {
    return this._userId;
  }
  public set userId(value: string) {
    this._userId = value;
  }

  public get listingId(): string {
    return this._listingId;
  }
  public set listingId(value: string) {
    this._listingId = value;
  }

  public get createdAt(): Date {
    return this._createdAt;
  }
  public set createdAt(value: Date) {
    this._createdAt = value;
  }

  public get updatedAt(): Date | undefined {
    return this._updatedAt;
  }
  public set updatedAt(value: Date | undefined) {
    this._updatedAt = value;
  }
}
