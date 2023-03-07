import { BookingResourceURLBuilder, RestAPIClient } from "../../client";
import { InternalServerError } from "../../error";

interface CreateBookingCommandConstructor {
  userId: string;
  spaceId: string;
  imageUrls: string[];
  description: string;
  startDate: string;
  endDate: string;
}

export class CreateBookingCommand {
  public readonly userId: string;
  public readonly spaceId: string;
  public readonly imageUrls: string[];
  public readonly description: string;
  public readonly startDate: string;
  public readonly endDate: string;

  public constructor({
    userId,
    spaceId,
    imageUrls,
    description,
    startDate,
    endDate,
  }: CreateBookingCommandConstructor) {
    this.userId = userId;
    this.spaceId = spaceId;
    this.imageUrls = imageUrls;
    this.description = description;
    this.startDate = startDate;
    this.endDate = endDate;
  }
}

export class CreateBookingUseCase {
  public async execute(
    command: CreateBookingCommand
  ): Promise<{ uid: string }> {
    const { userId, spaceId, description, imageUrls, startDate, endDate } =
      command;
    if (
      !userId ||
      !spaceId ||
      !imageUrls ||
      !description ||
      !startDate ||
      !endDate
    ) {
      throw new InternalServerError();
    }
    const client = new RestAPIClient();
    const response = await client.post<{ uid: string }, CreateBookingCommand>(
      BookingResourceURLBuilder.createBooking(),
      {
        userId,
        spaceId,
        imageUrls,
        description,
        startDate,
        endDate,
      }
    );
    return { uid: response.data.uid };
  }
}
