import { BookingRestClient } from "../../client";
import { InternalServerError } from "../../error";

interface CreateBookingCommandConstructor {
  userId: string;
  spaceId: string;
  startDate: string;
  endDate: string;
}

export class CreateBookingCommand {
  public readonly userId: string;
  public readonly spaceId: string;
  public readonly startDate: string;
  public readonly endDate: string;

  public constructor({
    userId,
    spaceId,
    startDate,
    endDate,
  }: CreateBookingCommandConstructor) {
    this.userId = userId;
    this.spaceId = spaceId;
    this.startDate = startDate;
    this.endDate = endDate;
  }
}

export class CreateBookingUseCase {
  public async execute(
    command: CreateBookingCommand
  ): Promise<{ uid: string }> {
    const { userId, spaceId, startDate, endDate } = command;
    if (!userId || !spaceId || !startDate || !endDate) {
      throw new InternalServerError();
    }
    const bookingClient = new BookingRestClient();
    return await bookingClient.createBooking(
      userId,
      spaceId,
      startDate,
      endDate
    );
  }
}
