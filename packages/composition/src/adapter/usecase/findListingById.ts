import { ListingRestClient } from "../../client/listing";

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
    return client.findListingById(listingId);
  }
}
