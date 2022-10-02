import { ListingServiceImpl } from "../../../../application/service/ListingService";
import { LoggerUtil } from "../../../../utils";

export const handler = async (event: any) => {
  const logger = new LoggerUtil("handler");
  logger.info(event, "handler()");
  const service = await ListingServiceImpl.create();
  try {
    for await (const record of event.Records) {
      const body = JSON.parse(record.body);
      const booking = JSON.parse(body.Message);
      await service.occupyStorageItem(booking);
    }
  } catch (error) {
    logger.error(error, "handler()");
  }
};
