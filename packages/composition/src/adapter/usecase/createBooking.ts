import { BookingRestClient } from "../../client";
import { InternalServerError } from "../../error";

interface CreateBookingCommandConstructor {
  userId: string;
  spaceId: string;
  imageUrls: string[];
}

export class CreateBookingCommand {
  public readonly userId: string;
  public readonly spaceId: string;
  public readonly imageUrls: string[];

  public constructor({
    userId,
    spaceId,
    imageUrls,
  }: CreateBookingCommandConstructor) {
    this.userId = userId;
    this.spaceId = spaceId;
    this.imageUrls = imageUrls;
  }
}

export class CreateBookingUseCase {
  public async execute(
    command: CreateBookingCommand
  ): Promise<{ uid: string }> {
    const { userId, spaceId, imageUrls } = command;
    if (!userId || !spaceId || !imageUrls) {
      throw new InternalServerError();
    }
    const bookingClient = new BookingRestClient();
    return await bookingClient.createBooking(userId, spaceId, imageUrls);
  }
}
