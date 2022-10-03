import { EmailAddress } from "../domain/model";
import { StorageItemInterface } from "./Item";

export interface UserInterface {
  id?: number;
  uid: string;
  firstName: string;
  lastName: string;
  emailAddress: string;
  password: string;
  createdAt?: string;
  updatedAt?: string;
  items: StorageItemInterface[];
}

export interface UserRawInterface {
  id: number;
  uid: string;
  first_name: string;
  last_name: string;
  email_address: string;
  password: string;
  created_at: string;
  updated_at?: string;
  items: StorageItemInterface[];
}

export interface UserConstructor {
  id?: number;
  uid?: string;
  firstName: string;
  lastName: string;
  emailAddress: EmailAddress;
  password: string;
  createdAt?: Date;
  updatedAt?: Date;
  items?: StorageItemInterface[];
}
