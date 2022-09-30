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
  public readonly value: string;

  public constructor(value: string) {
    if (!this.isValidStreetAddress(value)) {
      throw new Error("Invalid street address provided");
    }
    this.value = value;
  }

  private isValidStreetAddress(value: string) {
    return value.length > 0;
  }

  public get country(): string {
    return "Canada";
  }

  public get postalCode(): string {
    return this.value; // use regex or parsing to get the postal code
  }

  // have helper methods to get the portitions seperately like (country, postal code)
}
