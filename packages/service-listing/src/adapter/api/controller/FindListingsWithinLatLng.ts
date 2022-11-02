import { ListingServiceImpl } from "../../service";
import { LoggerUtil } from "../../../utils";

export const handler = async (event: any) => {
  const logger = new LoggerUtil("handler");
  console.log("FIND LISTINGS WITHIN LAT LNG EVENT: ", event);
  const { latitude, longitude, range } = event.arguments;
  logger.info(
    {
      latitude,
      longitude,
      range,
    },
    "handler"
  );
  const service = await ListingServiceImpl.create();
  return await service.findListingsWithinLatLng(latitude, longitude, range);
};
