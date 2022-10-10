import { Listing } from "../../domain/model";

export interface ListingRepository {
  setup(): Promise<void>;
  save(data: Listing): Promise<Listing>;
  delete(uid: string): Promise<void>;
  findOneById(uid: string): Promise<Listing>;
  findManyByLatLng(
    latitude: number,
    longitude: number,
    range: number
  ): Promise<Listing[]>;
}
