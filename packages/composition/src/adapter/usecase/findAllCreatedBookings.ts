import { BookingRestClient } from "../../client";
import { InternalServerError } from "../../error";

interface FindAllCreatedBookingsCommandConstructor {
  userId: string;
}

export class FindAllCreatedBookingsCommand {
  public readonly userId: string;

  public constructor({ userId }: FindAllCreatedBookingsCommandConstructor) {
    this.userId = userId;
  }
}

export class FindAllCreatedBookingsUseCase {
  public async execute(command: FindAllCreatedBookingsCommand) {
    const { userId } = command;
    if (!userId) {
      throw new InternalServerError();
    }
    const client = new BookingRestClient();
    const data = await client.findAllCreatedBookingsByUserId(userId);
    return data?.bookings;
  }
}
