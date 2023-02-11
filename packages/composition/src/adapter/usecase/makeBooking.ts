import { BookingRestClient } from "../../client";
import { InternalServerError } from "../../error";

interface MakeBookingCommandConstructor {
  userId: string;
  listingId: string;
  items: BookingItem[];
}

export class MakeBookingCommand {
  public readonly userId: string;
  public readonly listingId: string;
  public readonly items: BookingItem[];

  public constructor({
    userId,
    listingId,
    items,
  }: MakeBookingCommandConstructor) {
    this.userId = userId;
    this.listingId = listingId;
    this.items = items;
  }
}

export class MakeBookingUseCase {
  public async execute(command: MakeBookingCommand): Promise<{ uid: string }> {
    const { userId, listingId, items } = command;
    if (!userId || !listingId || !items) {
      throw new InternalServerError();
    }
    const bookingClient = new BookingRestClient();
    return await bookingClient.makeBooking(userId, listingId, items);
  }
}
