import { BookingServiceImpl } from "../../../app/service";

export const handler = async (event: any) => {
  const { amount, currency, userId, listingId, items } = event.arguments;
  const service = await BookingServiceImpl.create();
  return await service.makeBooking(amount, currency, userId, listingId, items);
};
