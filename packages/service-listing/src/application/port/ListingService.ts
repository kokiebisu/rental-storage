import { AggregatedListingInterface, ListingInterface } from "../../types";

export interface ListingService {
  findListingsWithinLatLng(
    latitude: number,
    longitude: number,
    range: number
  ): Promise<AggregatedListingInterface[]>;
  findListingById(id: number): Promise<AggregatedListingInterface | null>;
  addListing(data: ListingInterface): Promise<boolean>;
  removeListingById(id: number): Promise<boolean>;
}
