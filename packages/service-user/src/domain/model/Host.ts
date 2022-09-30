import { HostConstructor } from "../../types";
import { User } from "./User";

export class Host extends User {
  public constructor({
    id,
    firstName,
    lastName,
    createdAt,
    updatedAt,
  }: HostConstructor) {
    super({ id, firstName, lastName, createdAt, updatedAt });
  }
}
