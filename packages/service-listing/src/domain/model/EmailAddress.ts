export class EmailAddress {
  public readonly value: string;

  public constructor(value: string) {
    this.isValidEmailAddress(value);
    this.value = value;
  }

  private isValidEmailAddress(value: string) {
    if (!value) {
      throw new Error("Email address was not provided");
    }
    if (value.length < 5) {
      throw new Error("Email address should be longer");
    }
  }
}
