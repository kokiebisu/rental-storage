import { ListingRestClient } from "../../client";
import { InternalServerError } from "../../error";

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
  public async execute(
    command: RemoveListingByIdCommand
  ): Promise<{ uid: string }> {
    const { listingId } = command;
    if (!listingId) {
      throw new InternalServerError();
    }
    const listingClient = new ListingRestClient();
    return await listingClient.removeListingById(listingId);
  }
}
