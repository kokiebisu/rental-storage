import { BookingServiceImpl } from "../service";

export const makeBooking = async (event: any) => {
  const {} = event.arguments;
  const service = await BookingServiceImpl.create();
  return await service.makeBooking();
};
