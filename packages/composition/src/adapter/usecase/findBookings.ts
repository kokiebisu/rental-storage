import { BookingResourceURLBuilder, RestAPIClient } from "../../client";

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
  public async execute(command: FindBookingsCommand): Promise<IBooking[]> {
    const { spaceId } = command;
    const client = new RestAPIClient();
    const response = await client.get<{
      bookings: (Omit<IBooking, "uid"> & { uid: string })[];
    }>(BookingResourceURLBuilder.findAllBookings({ spaceId }));
    return response.data.bookings.map(
      (booking: Omit<IBooking, "uid"> & { uid: string }) => {
        return {
          ...booking,
          id: booking.uid,
        };
      }
    );
  }
}
