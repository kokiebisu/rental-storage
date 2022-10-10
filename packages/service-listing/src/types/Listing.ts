import { StreetAddress } from "../domain/model";

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

// export interface AggregatedListingInterface {
//   id?: number;
//   uid: string;
//   lender: LenderInterface;
//   streetAddress: string;
//   latitude: number;
//   longitude: number;
//   imageUrls: string[];
// }

export interface ListingInterface {
  id?: number;
  uid: string;
  lenderId: string;
  streetAddress: string;
  latitude: number;
  longitude: number;
  imageUrls: string[];
}
