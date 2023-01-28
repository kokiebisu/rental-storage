import { ListingRestClient } from "../../client";
import { InternalServerError } from "../../error";

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
    if (!listingId) {
      throw new InternalServerError();
    }
    const client = new ListingRestClient();
    return client.findListingById(listingId);
  }
}
