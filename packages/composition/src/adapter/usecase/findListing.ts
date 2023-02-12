import { ListingRestClient } from "../../client";

interface FindListingCommandConstructor {
  id: string;
}

export class FindListingCommand {
  public readonly id: string;

  public constructor({ id }: FindListingCommandConstructor) {
    this.id = id;
  }
}

export class FindListingUseCase {
  public async execute(command: FindListingCommand): Promise<Listing> {
    const { id } = command;
    const client = new ListingRestClient();
    const data = await client.findListing(id);
    return {
      id: data.listing.uid,
      feeType: data.listing.feeType,
      feeAmount: data.listing.feeAmount,
      feeCurrency: data.listing.feeCurrency,
      imageUrls: data.listing.imageUrls,
      latitude: data.listing.latitude,
      longitude: data.listing.longitude,
      lenderId: data.listing.lenderId,
      streetAddress: data.listing.streetAddress,
      title: data.listing.title,
    };
  }
}
