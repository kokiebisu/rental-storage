import { ListingServiceImpl } from "../../../app/service";
import { LoggerUtil } from "../../../utils";

export const handler = async (event: any) => {
  const logger = new LoggerUtil("handler");
  const { uid: lenderId } = event.identity.resolverContext;
  const { streetAddress, latitude, longitude, imageUrls, title, fee } =
    event.arguments;
  logger.info(
    {
      lenderId,
      streetAddress,
      latitude,
      longitude,
      imageUrls,
      title,
      fee,
    },
    "handler"
  );
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
