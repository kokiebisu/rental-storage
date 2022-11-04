import { CurrencyType, RentalFeeType } from "../../domain/enum";
import { ListingInterface } from "../../domain/types";

export interface AddListing {
  lenderId: string;
  streetAddress: string;
  latitude: number;
  longitude: number;
  imageUrls: string[];
  title: string;
  fee: {
    amount: number;
    currency: CurrencyType;
    type: RentalFeeType;
  };
}

export interface ListingService {
  findListingsWithinLatLng(
    latitude: number,
    longitude: number,
    range: number
  ): Promise<ListingInterface[]>;
  findListingById(uid: string): Promise<ListingInterface>;
  findListingsByUserId(userId: string): Promise<ListingInterface[]>;
  addListing(data: AddListing): Promise<boolean>;
  removeListingById(uid: string): Promise<boolean>;
}
