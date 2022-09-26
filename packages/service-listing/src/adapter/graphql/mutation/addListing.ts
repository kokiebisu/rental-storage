import { ListingServiceImpl } from "../../../service";

export const handler = async (event: any) => {
  const { hostId, streetAddress, latitude, longitude, items } = event.arguments;
  const service = await ListingServiceImpl.create();
  return await service.addListing({
    hostId,
    streetAddress,
    latitude,
    longitude,
    items,
  });
};
