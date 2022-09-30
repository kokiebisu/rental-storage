import { UserConstructor } from "./User";

export interface HostInterface {
  id?: string;
  firstName: string;
  lastName: string;
  createdAt: string;
  updatedAt?: string;
}

export interface HostConstructor extends UserConstructor {}

export interface HostRawInterface {
  id?: string;
  first_name: string;
  last_name: string;
  created_at: string;
  updated_at?: string;
}
