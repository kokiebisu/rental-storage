import { BookingResourceURLBuilder, RestAPIClient } from "../../client";

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
    const client = new RestAPIClient();
    const response = await client.get<{
      booking: Omit<IBooking, "id"> & { uid: string };
    }>(BookingResourceURLBuilder.findBooking(id));
    return {
      ...response.data.booking,
      id: response.data?.booking.uid,
    };
  }
}
