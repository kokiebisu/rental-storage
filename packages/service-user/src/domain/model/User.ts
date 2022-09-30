import { UserConstructor } from "../../types";

export abstract class User {
  public readonly id?: string;
  public readonly firstName: string;
  public readonly lastName: string;
  public readonly createdAt: Date;
  public readonly updatedAt?: Date;

  public constructor({
    id,
    firstName,
    lastName,
    createdAt = new Date(),
    updatedAt,
  }: UserConstructor) {
    this.id = id;
    this.firstName = firstName;
    this.lastName = lastName;
    this.createdAt = createdAt;
    this.updatedAt = updatedAt;
  }
}
