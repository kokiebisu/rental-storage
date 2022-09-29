import {
  LoggerService,
  StorageItemInterface,
} from "@rental-storage-project/common";
import { ItemServiceImpl } from "../../../../application/service";

export const handler = async (event: any) => {
  const logger = new LoggerService("handler");
  const service = await ItemServiceImpl.create();
  logger.info(event, "handler");
  try {
    for (const record of event.Records) {
      const items = mapItem(record.dynamodb.NewImage);

      for (const item of items) {
        // add item row to table
        await service.addItem(item);
      }
    }
  } catch (error) {
    console.log(error);
  }
};

const mapItem = (data: any): StorageItemInterface[] => {
  const { items } = data;

  // return {
  //   id: id.S,
  //   status: status.S,
  //   amount: {
  //     value: parseInt(amount.M.value.N),
  //     currency: amount.M.currency.S,
  //   },
  //   userId: user_id.S,
  //   listingId: listing_id.S,
  //   createdAt: created_at.S,
  //   ...(updated_at && { updatedAt: updated_at.S }),
  // };
  return items.map((item: any) => {
    return {
      name: item.name.S,
      imageUrls: item.imageUrls.map((imageUrl: any) => imageUrl.S),
      userId: item.user_id.S,
      listingId: item.listing_id.S,
    };
  });
};

// NewImage: {
//   amount: { M: [Object] },
//   listing_id: { S: '1' },
//   user_id: { S: '1' },
//   created_at: { S: '9/25/2022, 8:11:59 PM' },
//   id: { S: 'dd9dbcdf-8260-401e-b199-63d328bced18' },
//   status: { S: 'created' }
// },
