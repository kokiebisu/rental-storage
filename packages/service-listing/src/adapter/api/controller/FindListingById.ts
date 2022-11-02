import { ListingServiceImpl } from "../../service";
import { LoggerUtil } from "../../../utils";

export const handler = async (event: any) => {
  const logger = new LoggerUtil("handler");
  console.log("FIND LISTING BY ID EVENT: ", event);
  const { listingId } = event.pathParameters;
  logger.info(
    {
      listingId,
    },
    "handler"
  );
  const service = await ListingServiceImpl.create();
  return await service.findListingById(listingId);
};
