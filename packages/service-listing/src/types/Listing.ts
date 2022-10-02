import { StreetAddress } from "../domain/model";
import { LenderInterface } from "./Lender";
import { StorageItemInterface } from "./Item";

export interface ListingConstructor {
  id?: string;
  lenderId: number;
  streetAddress: StreetAddress;
  latitude: number;
  longitude: number;
  imageUrls: string[];
  items: string[];
}

export interface ListingRawInterface {
  id: string;
  lender_id: number;
  email_address: string;
  street_address: string;
  latitude: number;
  longitude: number;
  image_urls: string[];
  items: string[];
}

export interface AggregatedListingInterface {
  id?: string;
  lender: LenderInterface;
  streetAddress: string;
  latitude: number;
  longitude: number;
  imageUrls: string[];
  items: StorageItemInterface[];
}

export interface ListingInterface {
  id?: string;
  lenderId: number;
  streetAddress: string;
  latitude: number;
  longitude: number;
  imageUrls: string[];
  items: string[];
}
