import { ListingServiceImpl } from "../../../../../application/service/ListingService";

export const handler = async (event: any) => {
  const { lenderId, streetAddress, latitude, longitude, items, imageUrls } =
    event.arguments;
  const service = await ListingServiceImpl.create();
  return await service.addListing({
    lenderId,
    streetAddress,
    latitude,
    longitude,
    imageUrls,
  });
};
