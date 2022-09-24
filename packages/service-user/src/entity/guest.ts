import { User } from "./user";
import { GuestInterface } from "@rental-storage-project/common";
export interface GuestRawInterface {
  id: string;
  first_name: string;
  last_name: string;
}

export class Guest extends User {
  public constructor(firstName: string, lastName: string) {
    super(firstName, lastName);
  }

  public toDTO(): GuestInterface {
    return {
      firstName: this.firstName,
      lastName: this.lastName,
    };
  }
}
