import { ListingServiceImpl } from "../src/service";

export const addListing = async (event: any) => {
  const { hostId, streetAddress, latitude, longitude } = event.arguments;
  const service = await ListingServiceImpl.create();
  return await service.addListing({
    hostId,
    streetAddress,
    latitude,
    longitude,
  });
};

export const removeListingById = async (event: any) => {
  const { id } = event.arguments;
  const service = await ListingServiceImpl.create();
  return await service.removeListingById(id);
};
