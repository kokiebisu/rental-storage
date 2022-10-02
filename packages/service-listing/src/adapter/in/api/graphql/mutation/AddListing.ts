import { ListingServiceImpl } from "../../../../../application/service/ListingService";

export const handler = async (event: any) => {
  const { hostId, streetAddress, latitude, longitude, items, imageUrls } =
    event.arguments;
  const service = await ListingServiceImpl.create();
  return await service.addListing({
    hostId,
    streetAddress,
    latitude,
    longitude,
    imageUrls,
    items,
  });
};
