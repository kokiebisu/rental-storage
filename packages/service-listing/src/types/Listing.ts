import { Fee, StreetAddress } from "../Domain/Model";
import { FeeInterface } from "./Fee";

export interface ListingConstructor {
  id?: number;
  uid?: string;
  title: string;
  lenderId: string;
  streetAddress: StreetAddress;
  latitude: number;
  longitude: number;
  imageUrls: string[];
  fee: Fee;
}

export interface ListingRawInterface {
  id: number;
  uid: string;
  title: string;
  lender_id: string;
  email_address: string;
  street_address: string;
  latitude: number;
  longitude: number;
  image_urls: string[];
  fee_currency: string;
  fee_amount: number;
  fee_type: string;
}

export interface ListingInterface {
  id?: number;
  uid: string;
  title: string;
  lenderId: string;
  streetAddress: string;
  latitude: number;
  longitude: number;
  imageUrls: string[];
  fee: FeeInterface;
}
