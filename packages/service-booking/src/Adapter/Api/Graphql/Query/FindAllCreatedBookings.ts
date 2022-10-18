import { BookingServiceImpl } from "../../../../App/Service/BookingService";

export const handler = async (event: any) => {
  const {} = event.arguments;
  const service = await BookingServiceImpl.create();
  return await service.findAllCreatedBookings();
};
