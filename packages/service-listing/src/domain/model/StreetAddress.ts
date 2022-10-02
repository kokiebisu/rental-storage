export class City {
  public readonly value: string;

  public constructor(value: string) {
    this.isValidCity(value);
    this.value = value;
  }

  private isValidCity(value: string) {
    if (!value) {
      throw new Error("Unable to create City instance");
    }
  }
}

export class PostalCode {
  public readonly value: string;

  public constructor(value: string) {
    this.isValidPostalCode(value);
    this.value = value;
  }

  private isValidPostalCode(value: string) {
    if (!value) {
      throw new Error("Unable to create PostalCode instance");
    }
  }
}

export class StreetAddress {
  public readonly value: string;

  public constructor(value: string) {
    this.isValidStreetAddress(value);
    this.value = value;
  }

  private isValidStreetAddress(value: string) {
    if (!value) {
      throw new Error("Invalid street address provided");
    }
    if (value.length === 0) {
      throw new Error("Provided street address cannot be an empty string");
    }
  }

  public get country(): string {
    return "Canada";
  }

  public get postalCode(): string {
    return this.value; // use regex or parsing to get the postal code
  }

  // have helper methods to get the portitions seperately like (country, postal code)
}
