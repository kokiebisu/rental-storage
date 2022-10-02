import { ListingInterface } from "../../types";

export interface ListingRepository {
  setup(): Promise<void>;
  save(data: Omit<ListingInterface, "id">): Promise<void>;
  addItemToListing(
    id: number,
    itemId: number
  ): Promise<{ insertId: any } | undefined>;
  delete(id: number): Promise<void>;
  findOneById(id: number): Promise<ListingInterface>;
  findManyByLatLng(
    latitude: number,
    longitude: number,
    range: number
  ): Promise<ListingInterface[]>;
}
