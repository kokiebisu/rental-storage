import { ListingInterface } from "../../types";

export interface ListingEventSender {
  listingCreated(data: ListingInterface): Promise<void>;
}
