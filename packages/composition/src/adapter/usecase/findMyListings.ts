import { ListingRestClient } from "../../client/listing";

interface FindMyListingsCommandConstructor {
  userId: string;
}

export class FindMyListingsCommand {
  public readonly userId: string;

  public constructor({ userId }: FindMyListingsCommandConstructor) {
    this.userId = userId;
  }
}

export class FindMyListingsUseCase {
  public async execute(command: FindMyListingsCommand) {
    const { userId } = command;
    const listingClient = new ListingRestClient();
    const response = await listingClient.findListingsByUserId(userId);
    return response.data;
  }
}
