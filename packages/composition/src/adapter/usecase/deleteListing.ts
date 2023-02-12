import { ListingRestClient } from "../../client";
import { InternalServerError } from "../../error";

interface DeleteListingCommandConstructor {
  id: string;
}

export class DeleteListingCommand {
  public readonly id: string;

  public constructor({ id }: DeleteListingCommandConstructor) {
    this.id = id;
  }
}

export class DeleteListingUseCase {
  public async execute(
    command: DeleteListingCommand
  ): Promise<{ uid: string }> {
    const { id } = command;
    if (!id) {
      throw new InternalServerError();
    }
    const listingClient = new ListingRestClient();
    return await listingClient.deleteListing(id);
  }
}
