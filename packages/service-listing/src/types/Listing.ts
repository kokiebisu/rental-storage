import { StreetAddress } from "../domain/model";
import { HostInterface } from "./Host";
import { StorageItemInterface } from "./Item";

export interface ListingConstructor {
  id?: string;
  hostId: string;
  streetAddress: StreetAddress;
  latitude: number;
  longitude: number;
  items: string[];
}

export interface ListingRawInterface {
  id: string;
  host_id: string;
  email_address: string;
  street_address: string;
  latitude: number;
  longitude: number;
  items: string[];
}

export interface AggregatedListingInterface {
  id?: string;
  host: HostInterface;
  streetAddress: string;
  latitude: number;
  longitude: number;
  items: StorageItemInterface[];
}

export interface ListingInterface {
  id?: string;
  hostId: string;
  streetAddress: string;
  latitude: number;
  longitude: number;
  items: string[];
}
