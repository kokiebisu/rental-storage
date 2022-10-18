import { Listing } from "../../Domain/Model";

export interface ListingRepository {
  setup(): Promise<void>;
  save(data: Listing): Promise<Listing>;
  delete(data: Listing): Promise<void>;
  findOneById(uid: string): Promise<Listing>;
  findManyByLatLng(
    latitude: number,
    longitude: number,
    range: number
  ): Promise<Listing[]>;
  findManyByUserId(userId: string): Promise<Listing[]>;
}
