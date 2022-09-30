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
    if (!this.validateName(firstName)) {
      throw new Error("Provided first name is invalid");
    }
    if (!this.validateName(lastName)) {
      throw new Error("Provided last name is invalid");
    }
    this.id = id;
    this.firstName = firstName;
    this.lastName = lastName;
    this.createdAt = createdAt;
    this.updatedAt = updatedAt;
  }

  private validateName(name: string) {
    return name.length > 0 && new RegExp(/^[a-z ,.'-]+$/i).test(name);
  }
}
