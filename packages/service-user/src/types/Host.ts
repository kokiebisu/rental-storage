import { UserConstructor } from "./User";

export interface HostInterface {
  id?: number;
  uid?: string;
  emailAddress: string;
  password: string;
  firstName: string;
  lastName: string;
  createdAt: string;
  updatedAt?: string;
}

export interface HostConstructor extends UserConstructor {}

export interface HostRawInterface {
  id?: number;
  uid?: string;
  email_address: string;
  password: string;
  first_name: string;
  last_name: string;
  created_at: string;
  updated_at?: string;
}
