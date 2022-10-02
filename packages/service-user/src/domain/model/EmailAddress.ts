export class EmailAddress {
  public readonly value: string;

  public constructor(value: string) {
    this.validateEmailAddress(value);
    this.value = value;
  }

  private validateEmailAddress(value: string) {
    if (!value) {
      throw new Error("Email address not provided");
    }
    if (
      !String(value)
        .toLowerCase()
        .match(
          /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        )
    ) {
      throw new Error("Email address is invalid");
    }
  }
}
