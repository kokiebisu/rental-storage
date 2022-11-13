import { ListingRestClient } from "../../client/listing";

interface RemoveListingByIdCommandConstructor {
  listingId: string;
}

export class RemoveListingByIdCommand {
  public readonly listingId: string;

  public constructor({ listingId }: RemoveListingByIdCommandConstructor) {
    this.listingId = listingId;
  }
}

export class RemoveListingByIdUseCase {
  public async execute(command: RemoveListingByIdCommand) {
    const { listingId } = command;
    const listingClient = new ListingRestClient();
    return listingClient.removeListingById(listingId);
  }
}
