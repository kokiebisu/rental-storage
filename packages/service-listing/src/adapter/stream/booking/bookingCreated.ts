import { LoggerService } from "@rental-storage-project/common";
import { ListingServiceImpl } from "../../../service";

export const handler = async (event: any) => {
  const logger = new LoggerService("handler");
  const service = await ListingServiceImpl.create();
  logger.info(event, "handler");
  try {
    for (const record of event.Records) {
      const bookingDTO = mapBooking(record.dynamodb.NewImage);
      await service.occupyStorageItem(bookingDTO);
    }
  } catch (error) {
    console.log(error);
  }
};

const mapBooking = (data: any) => {
  const { amount, listing_id, user_id, created_at, id, status, updated_at } =
    data;

  return {
    id: id.S,
    status: status.S,
    amount: {
      value: parseInt(amount.M.value.N),
      currency: amount.M.currency.S,
    },
    userId: user_id.S,
    listingId: listing_id.S,
    createdAt: created_at.S,
    ...(updated_at && { updatedAt: updated_at.S }),
  };
};

// NewImage: {
//   amount: { M: [Object] },
//   listing_id: { S: '1' },
//   user_id: { S: '1' },
//   created_at: { S: '9/25/2022, 8:11:59 PM' },
//   id: { S: 'dd9dbcdf-8260-401e-b199-63d328bced18' },
//   status: { S: 'created' }
// },
