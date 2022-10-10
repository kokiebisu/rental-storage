import { StorageItemInterface, UserConstructor } from "../../types";
import { v4 as uuid } from "uuid";
import { EmailAddress } from "./EmailAddress";
import { Payment } from "./Payment";

export class User {
  public readonly id?: number;
  public readonly uid: string;
  public readonly firstName: string;
  public readonly lastName: string;
  public readonly emailAddress: EmailAddress;
  public readonly password: string;
  public readonly items: StorageItemInterface[];
  public payment?: Payment;
  public readonly createdAt: Date;
  public readonly updatedAt?: Date;

  public constructor({
    id,
    uid = uuid(),
    firstName,
    lastName,
    emailAddress,
    password,
    items = [],
    payment,
    createdAt = new Date(),
    updatedAt,
  }: UserConstructor) {
    this.validateFirstName(firstName);
    this.validateLastName(lastName);
    this.id = id;
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

  private validateFirstName(value: string) {
    if (!value) {
      throw new Error("First name is not provided");
    }
    if (value.length == 0) {
      throw new Error("First name cannot be empty");
    }
    // if (new RegExp(/^[a-z ,.'-]+$/i).test(value)) {
    //   throw new Error("Provided first name is invalid");
    // }
  }

  private validateLastName(value: string) {
    if (!value) {
      throw new Error("Last name is not provided");
    }
    if (value.length == 0) {
      throw new Error("Last name cannot be empty");
    }
    // if (new RegExp(/^[a-z ,.'-]+$/i).test(value)) {
    //   throw new Error("Provided last name is invalid");
    // }
  }

  public setPayment(payment: Payment) {
    this.payment = payment;
  }
}
