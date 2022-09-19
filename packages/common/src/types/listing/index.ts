import { HostInterface } from "../user";

export interface AggregatedListingInterface {
  id?: string;
  host: HostInterface;
  streetAddress: string;
  latitude: number;
  longitude: number;
}

export interface ListingInterface {
  id?: string;
  hostId: string;
  streetAddress: string;
  latitude: number;
  longitude: number;
}
