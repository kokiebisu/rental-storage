import { HostInterface } from "./Host";
import { StorageItemInterface } from "./Item";

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
