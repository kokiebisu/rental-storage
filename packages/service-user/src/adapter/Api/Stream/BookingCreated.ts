import { ItemServiceImpl } from "../../../Application/Service";
import { LoggerUtil } from "../../../Utils";
import { ItemMapper } from "../../Mapper";

export const handler = async (event: any) => {
  const logger = new LoggerUtil("handler");
  const itemService = await ItemServiceImpl.create();
  logger.info(event, "handler");
  try {
    for (const record of event.Records) {
      const items = ItemMapper.toDTOFromBookingStream(record.dynamodb.NewImage);

      for (const item of items) {
        // add item row to table
        await itemService.addItem(item);
      }
    }
  } catch (error) {
    logger.error(error, "handler()");
  }
};
