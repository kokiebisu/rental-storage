import { ListingConstructor } from "../../types";
import { StreetAddress } from "./StreetAddress";

export class Listing {
  public readonly id?: string;
  public readonly lenderId: number;
  public readonly streetAddress: StreetAddress;
  public readonly latitude: number;
  public readonly longitude: number;
  public readonly imageUrls: string[];
  public readonly items: string[];

  public constructor({
    id,
    lenderId,
    streetAddress,
    latitude,
    longitude,
    imageUrls,
    items = [],
  }: ListingConstructor) {
    this.validateLenderId(lenderId);
    this.validateLatitude(latitude);
    this.validateLongitude(longitude);

    this.id = id;
    this.lenderId = lenderId;
    this.streetAddress = streetAddress;
    this.latitude = latitude;
    this.longitude = longitude;
    this.imageUrls = imageUrls;
    this.items = items;
  }

  private validateLenderId(lenderId: number) {
    if (!lenderId) {
      throw new Error("lenderId was not provided");
    }
  }

  private validateLatitude(latitude: number) {
    if (!latitude) {
      throw new Error("latitude was not provided");
    }
    if (
      !new RegExp(
        /^(\+|-)?(?:90(?:(?:\.0{1,6})?)|(?:[0-9]|[1-8][0-9])(?:(?:\.[0-9]{1,6})?))$/
      ).test(latitude.toString())
    ) {
      throw new Error("Provided latitude doesn't exist");
    }
  }

  private validateLongitude(longitude: number) {
    if (!longitude) {
      throw new Error("longitude was not provided");
    }
    if (
      !new RegExp(
        /^(\+|-)?(?:180(?:(?:\.0{1,6})?)|(?:[0-9]|[1-9][0-9]|1[0-7][0-9])(?:(?:\.[0-9]{1,6})?))$/
      ).test(longitude.toString())
    ) {
      throw new Error("Provided latitude doesn't exist");
    }
  }
}
