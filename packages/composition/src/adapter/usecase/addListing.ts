import { ListingRestClient } from "../../client/listing";

interface AddListingCommandConstructor {
  lenderId: string;
  streetAddress: string;
  latitude: number;
  longitude: number;
  imageUrls: string[];
  title: string;
  fee: any;
}

export class AddListingCommand {
  public readonly lenderId: string;
  public readonly streetAddress: string;
  public readonly latitude: number;
  public readonly longitude: number;
  public readonly imageUrls: string[];
  public readonly title: string;
  public readonly fee: any;

  public constructor({
    lenderId,
    streetAddress,
    latitude,
    longitude,
    imageUrls,
    title,
    fee,
  }: AddListingCommandConstructor) {
    this.lenderId = lenderId;
    this.streetAddress = streetAddress;
    this.latitude = latitude;
    this.longitude = longitude;
    this.imageUrls = imageUrls;
    this.title = title;
    this.fee = fee;
  }
}

export class AddListingUseCase {
  public async execute(command: AddListingCommand) {
    const {
      lenderId,
      streetAddress,
      latitude,
      longitude,
      imageUrls,
      title,
      fee,
    } = command;
    const listingClient = new ListingRestClient();
    return listingClient.addListing(
      lenderId,
      streetAddress,
      latitude,
      longitude,
      imageUrls,
      title,
      fee
    );
  }
}
