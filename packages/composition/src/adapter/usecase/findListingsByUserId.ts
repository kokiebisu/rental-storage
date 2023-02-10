import { ListingRestClient } from "../../client";

interface FindListingsByUserIdCommandConstructor {
  userId: string;
}

export class FindListingsByUserIdCommand {
  public readonly userId: string;

  public constructor({ userId }: FindListingsByUserIdCommandConstructor) {
    this.userId = userId;
  }
}

export class FindListingsByUserIdUseCase {
  public async execute(command: FindListingsByUserIdCommand) {
    const { userId } = command;
    const client = new ListingRestClient();
    const data = await client.findListingsByUserId(userId);
    return data?.listings;
  }
}
