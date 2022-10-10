import { ProviderType } from "../domain/model";

export interface PaymentConstructor {
  id?: number;
  uid: string;
  userId: string;
  providerType: string;
  customerId: string;
}

export interface PaymentInterface {
  id?: number;
  uid: string;
  customerId: string;
  userId: string;
  providerType: ProviderType;
}
