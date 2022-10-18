import { v4 as uuid } from "uuid";

import { ListingConstructor } from "../../Types";
import { StreetAddress } from "./StreetAddress";

export class Listing {
  private _id?: number;
  public readonly uid: string;
  public readonly lenderId: string;
  public readonly streetAddress: StreetAddress;
  public readonly latitude: number;
  public readonly longitude: number;
  public readonly imageUrls: string[];

  public constructor({
    id,
    uid = uuid(),
    lenderId,
    streetAddress,
    latitude,
    longitude,
    imageUrls,
  }: ListingConstructor) {
    this.validateLenderId(lenderId);
    this.validateLatitude(latitude);
    this.validateLongitude(longitude);

    this._id = id;
    this.uid = uid;
    this.lenderId = lenderId;
    this.streetAddress = streetAddress;
    this.latitude = latitude;
    this.longitude = longitude;
    this.imageUrls = imageUrls;
  }

  public get id(): number {
    if (!this._id) {
      throw new Error("id cannot be retrieved since it doesn't exist");
    }
    return this._id;
  }

  public set id(value: number) {
    this._id = value;
  }

  private validateLenderId(lenderId: string) {
    if (!lenderId) {
      throw new Error("lenderId was not provided");
    }
  }

  private validateLatitude(latitude: number) {
    if (!latitude) {
      throw new Error("latitude was not provided");
    }
    if (!new RegExp(/([0-9.-]+).+?([0-9.-]+)/).test(latitude.toString())) {
      throw new Error("Provided latitude doesn't exist");
    }
  }

  private validateLongitude(longitude: number) {
    if (!longitude) {
      throw new Error("longitude was not provided");
    }
    if (!new RegExp(/([0-9.-]+).+?([0-9.-]+)/).test(longitude.toString())) {
      throw new Error("Provided latitude doesn't exist");
    }
  }
}
