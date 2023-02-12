import { ListingRestClient } from "../../client";

interface FindListingsCommandConstructor {
  userId: string;
}

export class FindListingsCommand {
  public readonly userId: string;

  public constructor({ userId }: FindListingsCommandConstructor) {
    this.userId = userId;
  }
}

export class FindListingsUseCase {
  public async execute(command: FindListingsCommand): Promise<Listing[]> {
    const { userId } = command;
    const client = new ListingRestClient();
    const data = await client.findListings(userId);
    return data.listings.map((listing: any) => {
      return {
        id: listing.uid,
        ...listing,
      };
    });
  }
}
