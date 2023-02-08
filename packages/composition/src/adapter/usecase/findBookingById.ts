import { BookingRestClient } from "../../client";

interface FindBookingByIdCommandConstructor {
  bookingId: string;
}

export class FindBookingByIdCommand {
  public readonly bookingId: string;

  public constructor({ bookingId }: FindBookingByIdCommandConstructor) {
    this.bookingId = bookingId;
  }
}

export class FindBookingByIdUseCase {
  public async execute(command: FindBookingByIdCommand) {
    const { bookingId } = command;
    const client = new BookingRestClient();
    const result = await client.findBookingById(bookingId);
    return result?.booking;
  }
}
