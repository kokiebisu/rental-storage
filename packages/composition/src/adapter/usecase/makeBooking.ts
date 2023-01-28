import { BookingRestClient } from "../../client";
import { InternalServerError } from "../../error";

interface MakeBookingCommandConstructor {
  userId: string;
  amount: number;
  currency: string;
  listingId: string;
  items: unknown;
}

export class MakeBookingCommand {
  public readonly userId: string;
  public readonly amount: number;
  public readonly currency: string;
  public readonly listingId: string;
  public readonly items: unknown;

  public constructor({
    userId,
    amount,
    currency,
    listingId,
    items,
  }: MakeBookingCommandConstructor) {
    this.userId = userId;
    this.amount = amount;
    this.currency = currency;
    this.listingId = listingId;
    this.items = items;
  }
}

export class MakeBookingUseCase {
  public async execute(command: MakeBookingCommand) {
    const { userId, amount, currency, listingId, items } = command;
    if (!userId || !amount || !currency || !listingId || !items) {
      throw new InternalServerError();
    }
    const bookingClient = new BookingRestClient();
    return await bookingClient.makeBooking(
      amount,
      currency,
      userId,
      listingId,
      items
    );
  }
}
