import {
  AggregatedListingInterface,
  ListingInterface,
} from "@rental-storage-project/common";

export interface ListingService {
  findListingsWithinLatLng(
    latitude: number,
    longitude: number,
    range: number
  ): Promise<AggregatedListingInterface[]>;
  findListingById(id: string): Promise<AggregatedListingInterface | null>;
  addListing(data: ListingInterface): Promise<boolean>;
  removeListingById(id: string): Promise<boolean>;
}
