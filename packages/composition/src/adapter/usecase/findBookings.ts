import { BookingResourceURLBuilder, RestAPIClient } from "../../client";

interface FindBookingsCommandConstructor {
  spaceId: string;
  bookingStatus: "PENDING" | "APPROVED";
}

export class FindBookingsCommand {
  public readonly spaceId: string;
  public readonly bookingStatus: "PENDING" | "APPROVED";

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
    if (bookingStatus === "PENDING") {
      response = await client.get<{
        bookings: (Omit<IBooking, "uid"> & { uid: string })[];
      }>(builder.findPendingBookings({ spaceId }));
    } else if (bookingStatus === "APPROVED") {
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
