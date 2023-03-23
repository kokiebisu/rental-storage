import { RestAPIClient } from "../../client";
import { InternalServerError } from "../../error";
import { BookingResourceURLBuilder } from "../../resource";

interface AcceptBookingCommandConstructor {
  id: string;
}

export class AcceptBookingCommand {
  public readonly id: string;

  public constructor({ id }: AcceptBookingCommandConstructor) {
    this.id = id;
  }
}

export class AcceptBookingUseCase {
  public async execute(command: AcceptBookingCommand): Promise<IBooking> {
    const { id } = command;
    if (!id) {
      throw new InternalServerError();
    }
    const client = new RestAPIClient();
    const builder = new BookingResourceURLBuilder();
    try {
      const response = await client.post<
        {
          booking: Omit<IBooking, "id"> & { uid: string };
        },
        AcceptBookingCommand
      >(builder.acceptBooking(id), {
        id,
      });
      return { ...response.data.booking, id: response.data?.booking.uid };
    } catch (error) {
      throw new InternalServerError();
    }
  }
}
