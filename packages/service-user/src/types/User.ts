import { EmailAddress } from "../domain/model";
import { Password } from "../domain/model/Password";

export interface UserConstructor {
  id?: number;
  uid?: string;
  firstName: string;
  lastName: string;
  emailAddress: EmailAddress;
  password: Password;
  createdAt?: Date;
  updatedAt?: Date;
}
