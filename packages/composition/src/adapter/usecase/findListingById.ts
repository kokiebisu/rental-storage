import { ListingRestClient } from "../../client";

interface FindListingByIdCommandConstructor {
  listingId: string;
}

export class FindListingByIdCommand {
  public readonly listingId: string;

  public constructor({ listingId }: FindListingByIdCommandConstructor) {
    this.listingId = listingId;
  }
}

export class FindListingByIdUseCase {
  public async execute(command: FindListingByIdCommand) {
    const { listingId } = command;
    const client = new ListingRestClient();
    const data = await client.findListingById(listingId);
    return data?.listing;
  }
}
