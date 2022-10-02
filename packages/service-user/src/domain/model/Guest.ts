import { GuestConstructor } from "../../types";
import { User } from "./User";

export class Guest extends User {
  public readonly items: string[];

  public constructor({
    id,
    uid,
    emailAddress,
    firstName,
    lastName,
    items = [],
    password,
    createdAt = new Date(),
    updatedAt,
  }: GuestConstructor) {
    super({
      id,
      uid,
      emailAddress,
      firstName,
      lastName,
      password,
      createdAt,
      updatedAt,
    });
    this.items = items;
  }
}
