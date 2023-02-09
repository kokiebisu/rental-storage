import { ListingRestClient } from "../../client";
import { InternalServerError } from "../../error";

interface AddListingCommandConstructor {
  lenderId: string;
  streetAddress: string;
  latitude: number;
  longitude: number;
  imageUrls: string[];
  title: string;
  feeAmount: number;
  feeCurrency: string;
  feeType: string;
}

export class AddListingCommand {
  public readonly lenderId: string;
  public readonly streetAddress: string;
  public readonly latitude: number;
  public readonly longitude: number;
  public readonly imageUrls: string[];
  public readonly title: string;
  public readonly feeAmount: number;
  public readonly feeCurrency: string;
  public readonly feeType: string;

  public constructor({
    lenderId,
    streetAddress,
    latitude,
    longitude,
    imageUrls,
    title,
    feeAmount,
    feeCurrency,
    feeType,
  }: AddListingCommandConstructor) {
    this.lenderId = lenderId;
    this.streetAddress = streetAddress;
    this.latitude = latitude;
    this.longitude = longitude;
    this.imageUrls = imageUrls;
    this.title = title;
    this.feeAmount = feeAmount;
    this.feeCurrency = feeCurrency;
    this.feeType = feeType;
  }
}

export class AddListingUseCase {
  public async execute(
    command: AddListingCommand
  ): Promise<{ uid: string } | undefined> {
    const {
      lenderId,
      streetAddress,
      latitude,
      longitude,
      imageUrls,
      title,
      feeAmount,
      feeCurrency,
      feeType,
    } = command;
    if (
      !lenderId ||
      !streetAddress ||
      !latitude ||
      !longitude ||
      !imageUrls ||
      !title ||
      !feeAmount ||
      !feeCurrency ||
      !feeType
    ) {
      throw new InternalServerError();
    }
    const listingClient = new ListingRestClient();
    return await listingClient.addListing(
      lenderId,
      streetAddress,
      latitude,
      longitude,
      imageUrls,
      title,
      feeType,
      feeAmount,
      feeCurrency
    );
  }
}
