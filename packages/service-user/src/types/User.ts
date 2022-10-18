import { EmailAddress, Name, Payment, ProviderType } from "../Domain/Model";
import { ItemInterface } from "./Item";

export interface UserInterface {
  id?: number;
  uid: string;
  firstName: string;
  lastName: string;
  emailAddress: string;
  password: string;
  payment?: {
    providerId: string;
    providerType: string;
  };
  items: ItemInterface[];
  createdAt?: string;
  updatedAt?: string;
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
  items: ItemInterface[];
  payment_id?: number;
  payment_provider_id?: string;
  payment_provider_type?: string;
}

export interface UserConstructor {
  id?: number;
  uid?: string;
  firstName: Name;
  lastName: Name;
  emailAddress: EmailAddress;
  payment?: Payment;
  password: string;
  createdAt?: Date;
  updatedAt?: Date;
  items?: ItemInterface[];
}
