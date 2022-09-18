export class City {
  public readonly value: string;

  public constructor(value: string) {
    if (!this.isValidCity()) {
      throw new Error("Unable to create City instance");
    }
    this.value = value;
  }

  private isValidCity() {
    return true;
  }
}

export class PostalCode {
  public readonly value: string;

  public constructor(value: string) {
    if (!this.isValidPostalCode()) {
      throw new Error("Unable to create PostalCode instance");
    }
    this.value = value;
  }

  private isValidPostalCode() {
    return true;
  }
}

export class StreetAddress {
  private _country: string;
  private _postalCode: PostalCode;
  private _city: City;
  private _address: string;
  private _latitude: number;
  private _longitude: number;

  public constructor(
    country: string,
    postalCode: PostalCode,
    city: City,
    address: string,
    latitude: number,
    longitude: number
  ) {
    this._country = country;
    this._postalCode = postalCode;
    this._city = city;
    this._address = address;
    this._latitude = latitude;
    this._longitude = longitude;
  }

  public get country(): string {
    return this._country;
  }

  public get postalCode(): string {
    return this._postalCode.value;
  }

  public get city(): string {
    return this._city.value;
  }

  public get address(): string {
    return this._address;
  }

  public get latitude(): number {
    return this._latitude;
  }

  public get longitude(): number {
    return this._longitude;
  }
}
