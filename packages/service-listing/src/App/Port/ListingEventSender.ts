import { ListingInterface } from "../../Types";

export interface ListingEventSender {
  listingCreated(data: ListingInterface): Promise<void>;
}
