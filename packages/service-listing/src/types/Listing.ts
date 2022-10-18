import { StreetAddress } from "../Domain/Model";

export interface ListingConstructor {
  id?: number;
  uid?: string;
  lenderId: string;
  streetAddress: StreetAddress;
  latitude: number;
  longitude: number;
  imageUrls: string[];
}

export interface ListingRawInterface {
  id: number;
  uid: string;
  lender_id: string;
  email_address: string;
  street_address: string;
  latitude: number;
  longitude: number;
  image_urls: string[];
}

export interface ListingInterface {
  id?: number;
  uid: string;
  lenderId: string;
  streetAddress: string;
  latitude: number;
  longitude: number;
  imageUrls: string[];
}
