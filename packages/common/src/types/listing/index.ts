import { HostInterface } from "../user";

export interface AggregatedListingInterface {
  id?: string;
  host: HostInterface;
  emailAddress: string;
  streetAddress: string;
  latitude: number;
  longitude: number;
}

export interface ListingInterface {
  id?: string;
  hostId: string;
  emailAddress: string;
  streetAddress: string;
  latitude: number;
  longitude: number;
}
