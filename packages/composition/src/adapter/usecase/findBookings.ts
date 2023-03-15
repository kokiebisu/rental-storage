import { RestAPIClient } from "../../client";
import { BookingResourceURLBuilder } from "../../resource";

interface FindBookingsCommandConstructor {
  spaceId: string;
  bookingStatus: "pending" | "approved";
}

export class FindBookingsCommand {
  public readonly spaceId: string;
  public readonly bookingStatus: "pending" | "approved";

  public constructor({
    spaceId,
    bookingStatus,
  }: FindBookingsCommandConstructor) {
    this.spaceId = spaceId;
    this.bookingStatus = bookingStatus;
  }
}

export class FindBookingsUseCase {
  public async execute(command: FindBookingsCommand): Promise<IBooking[]> {
    const { spaceId, bookingStatus } = command;
    const client = new RestAPIClient();
    const builder = new BookingResourceURLBuilder();
    let response;
    if (bookingStatus === "pending") {
      response = await client.get<{
        bookings: (Omit<IBooking, "uid"> & { uid: string })[];
      }>(builder.findPendingBookings({ spaceId }));
    } else if (bookingStatus === "approved") {
      response = await client.get<{
        bookings: (Omit<IBooking, "uid"> & { uid: string })[];
      }>(builder.findApprovedBookings({ spaceId }));
    } else {
      throw new Error("Invalid booking status");
    }
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
