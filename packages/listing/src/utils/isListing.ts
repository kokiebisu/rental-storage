import { Listing, ListingInterface } from "../entity";

export const isListing = (data: any): data is Listing => {
  return (
    (data as ListingInterface).id !== undefined &&
    (data as ListingInterface).emailAddress !== undefined &&
    (data as ListingInterface).host !== undefined &&
    (data as ListingInterface).streetAddress !== undefined &&
  );
};
