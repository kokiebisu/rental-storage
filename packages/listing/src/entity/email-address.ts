export class EmailAddress {
  public readonly value: string;

  public constructor(value: string) {
    if (!this.isValid(value)) {
      throw new Error("Provided email address is invalid");
    }
    this.value = value;
  }

  private isValid(value: string) {
    return value.length > 3;
  }
}
