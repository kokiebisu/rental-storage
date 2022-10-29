import { ListingInterface } from "../../domain/types";

export interface ListingEventSender {
  listingCreated(data: ListingInterface): Promise<void>;
}
