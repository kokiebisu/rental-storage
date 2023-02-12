import { BookingRestClient } from "../../client";

interface FindBookingsCommandConstructor {
  spaceId: string;
}

export class FindBookingsCommand {
  public readonly spaceId: string;

  public constructor({ spaceId }: FindBookingsCommandConstructor) {
    this.spaceId = spaceId;
  }
}

export class FindBookingsUseCase {
  public async execute(command: FindBookingsCommand): Promise<Booking[]> {
    const { spaceId } = command;
    const client = new BookingRestClient();
    const result = await client.findBookings(spaceId);
    return result.listings.map((listing: any) => {
      return {
        ...listing,
        id: result?.booking.uid,
      };
    });
  }
}
