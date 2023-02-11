import { BookingRestClient } from "../../client";

interface FindBookingCommandConstructor {
  id: string;
}

export class FindBookingCommand {
  public readonly id: string;

  public constructor({ id }: FindBookingCommandConstructor) {
    this.id = id;
  }
}

export class FindBookingUseCase {
  public async execute(command: FindBookingCommand): Promise<Booking> {
    const { id } = command;
    const client = new BookingRestClient();
    const result = await client.findBooking(id);
    return {
      id: result?.booking.uid,
      status: result?.booking.status,
      userId: result?.booking.userId,
      listingId: result?.booking.listingId,
      items: result?.booking.items,
      createdAt: result?.booking.createdAt,
      updatedAt: result?.booking.updatedAt,
    };
  }
}
