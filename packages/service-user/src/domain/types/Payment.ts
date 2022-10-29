import { ProviderType } from "../model";

export interface PaymentConstructor {
  id?: number;
  userId: number;
  providerType: string;
  providerId: string;
}

export interface PaymentRawInterface {
  id: number;
  customer_id: string;
  user_id: string;
  provider_type: string;
}

export interface PaymentInterface {
  id?: number;
  providerId: string;
  userId: number;
  providerType: ProviderType;
}
