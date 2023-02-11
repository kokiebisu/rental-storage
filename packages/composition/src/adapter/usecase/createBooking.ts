import { BookingRestClient } from "../../client";
import { InternalServerError } from "../../error";

interface CreateBookingCommandConstructor {
  userId: string;
  listingId: string;
  items: BookingItem[];
}

export class CreateBookingCommand {
  public readonly userId: string;
  public readonly listingId: string;
  public readonly items: BookingItem[];

  public constructor({
    userId,
    listingId,
    items,
  }: CreateBookingCommandConstructor) {
    this.userId = userId;
    this.listingId = listingId;
    this.items = items;
  }
}

export class CreateBookingUseCase {
  public async execute(
    command: CreateBookingCommand
  ): Promise<{ uid: string }> {
    const { userId, listingId, items } = command;
    if (!userId || !listingId || !items) {
      throw new InternalServerError();
    }
    const bookingClient = new BookingRestClient();
    return await bookingClient.createBooking(userId, listingId, items);
  }
}
