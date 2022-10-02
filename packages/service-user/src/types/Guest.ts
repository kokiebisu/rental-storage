import { UserConstructor } from "./User";

export interface GuestInterface {
  id?: number;
  uid?: string;
  emailAddress: string;
  password: string;
  firstName: string;
  lastName: string;
  items: string[];
  createdAt: string;
  updatedAt?: string;
}

export interface GuestRawInterface {
  id: number;
  uid: string;
  email_address: string;
  password: string;
  first_name: string;
  last_name: string;
  items: string[];
  created_at: string;
  updated_at?: string;
}

export interface GuestConstructor extends UserConstructor {
  items?: string[];
}
