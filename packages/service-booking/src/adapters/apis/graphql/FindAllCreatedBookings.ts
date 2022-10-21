import { BookingServiceImpl } from "../../../apps/service";

export const handler = async (event: any) => {
  const {} = event.arguments;
  const service = await BookingServiceImpl.create();
  return await service.findAllCreatedBookings();
};
