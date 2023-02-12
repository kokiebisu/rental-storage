import { BookingRestClient } from "../../client";
import { InternalServerError } from "../../error";

interface CreateBookingCommandConstructor {
  userId: string;
  spaceId: string;
  items: BookingItem[];
}

export class CreateBookingCommand {
  public readonly userId: string;
  public readonly spaceId: string;
  public readonly items: BookingItem[];

  public constructor({
    userId,
    spaceId,
    items,
  }: CreateBookingCommandConstructor) {
    this.userId = userId;
    this.spaceId = spaceId;
    this.items = items;
  }
}

export class CreateBookingUseCase {
  public async execute(
    command: CreateBookingCommand
  ): Promise<{ uid: string }> {
    const { userId, spaceId, items } = command;
    if (!userId || !spaceId || !items) {
      throw new InternalServerError();
    }
    const bookingClient = new BookingRestClient();
    return await bookingClient.createBooking(userId, spaceId, items);
  }
}
