import { ListingInterface, StorageItemInterface } from "../../types";

export interface ListingRepository {
  setup(): Promise<void>;
  save(data: Omit<ListingInterface, "id">): Promise<void>;
  addItemToListing(
    listingId: string,
    itemId: StorageItemInterface[]
  ): Promise<any>;
  delete(id: string): Promise<void>;
  findOneById(listingId: string): Promise<ListingInterface>;
  findManyByLatLng(
    latitude: number,
    longitude: number,
    range: number
  ): Promise<ListingInterface[]>;
}
