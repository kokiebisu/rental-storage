import { RestAPIClient } from "../../client";
import { InternalServerError } from "../../error";
import {
  BookingResourceURLBuilder,
  UserResourceURLBuilder,
} from "../../resource";

interface FindUserCommandConstructor {
  id: string;
}

export class FindUserCommand {
  public readonly id: string;

  public constructor({ id }: FindUserCommandConstructor) {
    this.id = id;
  }
}

export class FindUserUseCase {
  public async execute(command: FindUserCommand): Promise<IUser> {
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
    const bookingApprovedResponse = await client.get<{
      bookings: (Omit<IBooking, "id"> & { uid: string })[];
    }>(bookingResourceBuilder.findApprovedBookings({ userId: id }));
    const bookingPendingResponse = await client.get<{
      bookings: (Omit<IBooking, "id"> & { uid: string })[];
    }>(bookingResourceBuilder.findPendingBookings({ userId: id }));
    return {
      ...userResponse.data.user,
      id: userResponse.data.user.uid,
      bookings: {
        pending: bookingPendingResponse.data.bookings.map((booking) => {
          return {
            ...booking,
            id: booking.uid,
          };
        }),
        approved: bookingApprovedResponse.data.bookings.map((booking) => {
          return {
            ...booking,
            id: booking.uid,
          };
        }),
      },
    };
  }
}
