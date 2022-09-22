import { AmountInterface } from "../amount";
import { ListingInterface } from "../listing";
import { GuestInterface } from "../user";

export interface AggregatedBookingInterface {
  id?: string;
  status: string;
  amount: AmountInterface;
  user: GuestInterface;
  listing: ListingInterface;
  createdAt: Date;
  updatedAt?: Date;
}

export interface BookingInterface {
  id?: string;
  status: string;
  amount: AmountInterface;
  userId: string;
  listingId: string;
  createdAt: Date;
  updatedAt?: Date;
}
