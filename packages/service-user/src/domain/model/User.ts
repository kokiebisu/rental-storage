import { ItemInterface, UserConstructor } from "../../Types";
import { v4 as uuid } from "uuid";
import { EmailAddress } from "./EmailAddress";
import { Payment } from "./Payment";
import { Name } from "./Name";

export class User {
  private _id?: number;
  public readonly uid: string;
  public readonly firstName: Name;
  public readonly lastName: Name;
  public readonly emailAddress: EmailAddress;
  public readonly password: string;
  public readonly items: ItemInterface[];
  public payment?: Payment;
  public readonly createdAt: Date;
  public readonly updatedAt?: Date;

  public constructor({
    id,
    uid = uuid(),
    firstName,
    lastName,
    emailAddress,
    items = [],
    password,
    payment,
    createdAt = new Date(),
    updatedAt,
  }: UserConstructor) {
    this._id = id;
    this.uid = uid;
    this.firstName = firstName;
    this.lastName = lastName;
    this.emailAddress = emailAddress;
    this.password = password;
    this.items = items;
    this.payment = payment;
    this.createdAt = createdAt;
    this.updatedAt = updatedAt;
  }

  public get id(): number {
    if (!this._id) {
      throw new Error("id doesn't exist");
    }
    return this._id;
  }

  public set id(value: number) {
    this.validateId(value);
    this._id = value;
  }

  private validateId(value: number) {
    if (!value) {
      throw new Error("Provided id is not valid");
    }
    if (isNaN(value)) {
      throw new Error("Provided id is not NaN");
    }
  }

  public setPayment(payment: Payment) {
    this.payment = payment;
  }

  public static isUser(data: User | string): data is User {
    return (data as User).emailAddress !== undefined;
  }
}
