import { UserConstructor } from "./User";

export interface GuestInterface {
  id?: string;
  firstName: string;
  lastName: string;
  items: string[];
}

export interface GuestRawInterface {
  id: string;
  first_name: string;
  last_name: string;
  items: string[];
}

export interface GuestConstructor extends UserConstructor {
  items?: string[];
}
