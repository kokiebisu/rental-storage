import { StreetAddress } from "./street-address";
import { EmailAddress } from "./email-address";

export interface ListingRawInterface {
  id: string;
  host_id: string;
  email_address: string;
  street_address: string;
  latitude: number;
  longitude: number;
}

export class Listing {
  private _id?: string;
  public readonly hostId: string;
  private _emailAddress: EmailAddress;
  private _streetAddress: StreetAddress;

  public readonly latitude: number;
  public readonly longitude: number;

  public constructor(
    hostId: string,
    emailAddress: EmailAddress,
    streetAddress: StreetAddress,
    latitude: number,
    longitude: number
  ) {
    this.hostId = hostId;
    this._emailAddress = emailAddress;
    this._streetAddress = streetAddress;
    this.latitude = latitude;
    this.longitude = longitude;
  }

  public get id(): string | undefined {
    if (!this._id) {
      throw new Error("id doesn't exist for listing");
    }
    return this._id;
  }

  public set id(value: string | undefined) {
    this._id = value;
  }

  public get emailAddress(): string {
    return this._emailAddress.value;
  }

  public get streetAddress(): string {
    return this._streetAddress.value;
  }
}
