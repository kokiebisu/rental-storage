import { UserConstructor } from "./User";

export interface GuestInterface {
  id?: string;
  firstName: string;
  lastName: string;
  items: string[];
  createdAt: string;
  updatedAt?: string;
}

export interface GuestRawInterface {
  id: string;
  first_name: string;
  last_name: string;
  items: string[];
  created_at: string;
  updated_at?: string;
}

export interface GuestConstructor extends UserConstructor {
  items?: string[];
}
