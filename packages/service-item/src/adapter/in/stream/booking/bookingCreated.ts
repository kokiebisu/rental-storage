import { ItemServiceImpl } from "../../../../application/service";
import { LoggerUtil } from "../../../../utils";
import { StorageItemMapper } from "../../mapper";

export const handler = async (event: any) => {
  const logger = new LoggerUtil("handler");
  const service = await ItemServiceImpl.create();
  logger.info(event, "handler");
  try {
    for (const record of event.Records) {
      const items = StorageItemMapper.toDTOFromBookingStream(
        record.dynamodb.NewImage
      );

      for (const item of items) {
        // add item row to table
        await service.addItem(item);
      }
    }
  } catch (error) {
    logger.error(error, "handler()");
  }
};
