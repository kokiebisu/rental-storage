import { BookingRestClient } from "../../client";
import { InternalServerError } from "../../error";

interface CreateBookingCommandConstructor {
  userId: string;
  spaceId: string;
}

export class CreateBookingCommand {
  public readonly userId: string;
  public readonly spaceId: string;

  public constructor({ userId, spaceId }: CreateBookingCommandConstructor) {
    this.userId = userId;
    this.spaceId = spaceId;
  }
}

export class CreateBookingUseCase {
  public async execute(
    command: CreateBookingCommand
  ): Promise<{ uid: string }> {
    const { userId, spaceId } = command;
    if (!userId || !spaceId) {
      throw new InternalServerError();
    }
    const bookingClient = new BookingRestClient();
    return await bookingClient.createBooking(userId, spaceId);
  }
}
