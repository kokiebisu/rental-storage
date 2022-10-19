import { ListingServiceImpl } from "../../../app/service";

export const handler = async (event: any) => {
  const { uid: lenderId } = event.identity.resolverContext;
  const { streetAddress, latitude, longitude, imageUrls, title, fee } =
    event.arguments;
  const service = await ListingServiceImpl.create();
  return await service.addListing({
    lenderId,
    streetAddress,
    latitude,
    longitude,
    imageUrls,
    title,
    fee,
  });
};
