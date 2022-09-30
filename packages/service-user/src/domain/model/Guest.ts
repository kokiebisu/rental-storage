import { GuestConstructor } from "../../types";
import { User } from "./User";

export class Guest extends User {
  public readonly items: string[];

  public constructor({
    id,
    firstName,
    lastName,
    items = [],
    createdAt = new Date(),
    updatedAt,
  }: GuestConstructor) {
    super({ id, firstName, lastName, createdAt, updatedAt });
    this.items = items;
  }
}
