import { EmailAddress, Payment, ProviderType } from "../domain/model";
import { StorageItemInterface } from "./Item";
import { PaymentInterface } from "./Payment";

export interface UserInterface {
  id?: number;
  uid: string;
  firstName: string;
  lastName: string;
  emailAddress: string;
  password: string;
  payment?: {
    id?: string;
    providerType: ProviderType;
  };
  items: StorageItemInterface[];
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
  items: StorageItemInterface[];
  payment_id: string;
  provider_type: ProviderType;
}

export interface UserConstructor {
  id?: number;
  uid?: string;
  firstName: string;
  lastName: string;
  emailAddress: EmailAddress;
  payment?: Payment;
  password: string;
  createdAt?: Date;
  updatedAt?: Date;
  items?: StorageItemInterface[];
}
