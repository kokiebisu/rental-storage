import { ListingInterface } from "../entity";

export const isListing = (data: any): data is ListingInterface => {
  return (
    (data as ListingInterface).hostId !== undefined &&
    (data as ListingInterface).emailAddress !== undefined &&
    (data as ListingInterface).streetAddress !== undefined &&
    (data as ListingInterface).latitude !== undefined &&
    (data as ListingInterface).longitude !== undefined
  );
};
