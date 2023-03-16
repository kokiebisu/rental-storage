import { RestAPIClient } from "../../client";
import {
  BookingResourceURLBuilder,
  SpaceResourceURLBuilder,
} from "../../resource";

interface FindSpacesCommandConstructor {
  userId?: string;
  offset?: number;
  limit?: number;
}

export class FindSpacesCommand {
  public readonly userId?: string;
  public readonly offset?: number;
  public readonly limit?: number;

  public constructor({ userId, offset, limit }: FindSpacesCommandConstructor) {
    this.userId = userId;
    this.offset = offset;
    this.limit = limit;
  }
}

export class FindSpacesUseCase {
  public async execute(command: FindSpacesCommand): Promise<ISpace[]> {
    const { userId, offset, limit } = command;
    const client = new RestAPIClient();
    const builder = new SpaceResourceURLBuilder();
    const param: {
      userId?: string;
      offset?: number;
      limit?: number;
    } = {};
    if (userId) {
      param["userId"] = userId;
    }
    if (offset !== undefined) {
      param["offset"] = offset;
    }
    if (limit !== undefined) {
      param["limit"] = limit;
    }
    const url = builder.findSpaces(param);

    const spaceResponse = await client.get<{
      spaces: (Omit<ISpace, "id"> & { uid: string })[];
    }>(url);

    const bookingResourceBuilder = new BookingResourceURLBuilder();

    const pendingBookings: IBooking[] = [];

    for (const space of spaceResponse.data.spaces) {
      const bookingPendingResponse = await client.get<{
        bookings: (Omit<IBooking, "id"> & { uid: string })[];
      }>(bookingResourceBuilder.findPendingBookings({ spaceId: space.uid }));
      for (const booking of bookingPendingResponse.data.bookings) {
        pendingBookings.push({
          ...booking,
          id: booking.uid,
        });
      }
    }

    const approvedBookings: IBooking[] = [];

    for (const space of spaceResponse.data.spaces) {
      const bookingApprovedResponse = await client.get<{
        bookings: (Omit<IBooking, "id"> & { uid: string })[];
      }>(bookingResourceBuilder.findApprovedBookings({ spaceId: space.uid }));
      for (const booking of bookingApprovedResponse.data.bookings) {
        approvedBookings.push({
          ...booking,
          id: booking.uid,
        });
      }
    }

    return spaceResponse.data.spaces.map(
      (space: Omit<ISpace, "id"> & { uid: string }) => {
        return {
          id: space.uid,
          ...space,
          bookings: {
            pending: pendingBookings,
            approved: approvedBookings,
          },
        };
      }
    );
  }
}
