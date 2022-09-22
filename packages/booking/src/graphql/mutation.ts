import { BookingServiceImpl } from "../service";

export const makeBooking = async (event: any) => {
  const { amount, currency, userId, listingId } = event.arguments;
  const service = await BookingServiceImpl.create();
  return await service.makeBooking(amount, currency, userId, listingId);
};
