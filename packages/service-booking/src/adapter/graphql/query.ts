import { BookingServiceImpl } from "../../service";

export const findAllCreatedBookings = async (event: any) => {
  const {} = event.arguments;
  const service = await BookingServiceImpl.create();
  return await service.findAllCreatedBookings();
};
