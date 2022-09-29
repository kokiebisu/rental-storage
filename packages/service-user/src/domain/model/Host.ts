import { User } from "./User";
import { HostInterface } from "@rental-storage-project/common";

export interface HostRawInterface {
  id?: string;
  first_name: string;
  last_name: string;
}

export class Host extends User {
  public constructor(firstName: string, lastName: string) {
    super(firstName, lastName);
  }

  public toDTO(): HostInterface {
    return {
      firstName: this.firstName,
      lastName: this.lastName,
    };
  }
}
