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
  public async execute(command: FindBookingCommand): Promise<IBooking> {
    const { id } = command;
    const client = new BookingRestClient();
    const result = await client.findBooking(id);
    return {
      ...result.booking,
      id: result?.booking.uid,
    };
  }
}
