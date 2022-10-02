import { HostConstructor } from "../../types";
import { User } from "./User";

export class Host extends User {
  public constructor({
    id,
    emailAddress,
    firstName,
    lastName,
    password,
    createdAt,
    updatedAt,
  }: HostConstructor) {
    super({
      id,
      emailAddress,
      firstName,
      lastName,
      password,
      createdAt,
      updatedAt,
    });
  }
}
