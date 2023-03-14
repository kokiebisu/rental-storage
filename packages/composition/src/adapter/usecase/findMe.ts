import { RestAPIClient } from "../../client";
import { InternalServerError } from "../../error";
import {
  BookingResourceURLBuilder,
  UserResourceURLBuilder,
} from "../../resource";

interface FindMeCommandConstructor {
  id: string;
}

export class FindMeCommand {
  public readonly id: string;

  public constructor({ id }: FindMeCommandConstructor) {
    this.id = id;
  }
}

export class FindMeUseCase {
  public async execute(command: FindMeCommand): Promise<IUser> {
    const { id } = command;
    if (!id) {
      throw new InternalServerError();
    }

    const client = new RestAPIClient();
    const userResourceBuilder = new UserResourceURLBuilder();
    const userResponse = await client.get<{
      user: Omit<IUser, "id"> & { uid: string };
    }>(userResourceBuilder.findUser(id));

    const bookingResourceBuilder = new BookingResourceURLBuilder();
    const pendingBookingsResponse = await client.get<{
      bookings: (Omit<IBooking, "id"> & { uid: string })[];
    }>(
      bookingResourceBuilder.findPendingBookings({
        userId: userResponse.data.user.uid,
      })
    );
    const approvedBookingsResponse = await client.get<{
      bookings: (Omit<IBooking, "id"> & { uid: string })[];
    }>(
      bookingResourceBuilder.findApprovedBookings({
        userId: userResponse.data.user.uid,
      })
    );
    return {
      id: userResponse.data.user.uid,
      firstName: userResponse.data.user.firstName,
      lastName: userResponse.data.user.lastName,
      emailAddress: userResponse.data.user.emailAddress,
      streetAddress: userResponse.data.user.streetAddress,
      createdAt: userResponse.data.user.createdAt,
      updatedAt: userResponse.data.user.updatedAt,
      bookings: {
        pending: pendingBookingsResponse.data.bookings.map((booking) => {
          return {
            ...booking,
            id: booking.uid,
          };
        }),
        approved: approvedBookingsResponse.data.bookings.map((booking) => {
          return {
            ...booking,
            id: booking.uid,
          };
        }),
      },
    };
  }
}
