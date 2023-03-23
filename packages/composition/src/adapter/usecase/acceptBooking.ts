import { RestAPIClient } from "../../client";
import { InternalServerError } from "../../error";
import { BookingResourceURLBuilder } from "../../resource";

interface AcceptBookingCommandConstructor {
  bookingId: string;
}

export class AcceptBookingCommand {
  public readonly bookingId: string;

  public constructor({ bookingId }: AcceptBookingCommandConstructor) {
    this.bookingId = bookingId;
  }
}

export class AcceptBookingUseCase {
  public async execute(command: AcceptBookingCommand): Promise<IBooking> {
    const { bookingId } = command;
    if (!bookingId) {
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
      >(builder.acceptBooking(bookingId), {
        bookingId,
      });
      return { ...response.data.booking, id: response.data?.booking.uid };
    } catch (error) {
      throw new InternalServerError();
    }
  }
}
