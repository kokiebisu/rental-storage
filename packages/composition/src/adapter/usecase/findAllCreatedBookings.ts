import { BookingRestClient } from "../../client/booking";

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
    const client = new BookingRestClient();
    return await client.findAllCreatedBookingsByUserId(userId);
  }
}
