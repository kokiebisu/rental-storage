import { StreetAddress } from "./StreetAddress";

export interface ListingRawInterface {
  id: string;
  host_id: string;
  email_address: string;
  street_address: string;
  latitude: number;
  longitude: number;
  items: string[];
}

export class Listing {
  private _id?: string;
  public readonly hostId: string;
  private _streetAddress: StreetAddress;

  public readonly latitude: number;
  public readonly longitude: number;
  public readonly items: string[];

  public constructor(
    hostId: string,
    streetAddress: StreetAddress,
    latitude: number,
    longitude: number,
    items: string[] = []
  ) {
    this.hostId = hostId;
    this._streetAddress = streetAddress;
    this.latitude = latitude;
    this.longitude = longitude;
    this.items = items;
  }

  public get id(): string | undefined {
    return this._id;
  }

  public set id(value: string | undefined) {
    this._id = value;
  }

  public get streetAddress(): string {
    return this._streetAddress.value;
  }
}
