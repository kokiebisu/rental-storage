export abstract class User {
  private _id?: string;
  public readonly firstName: string;
  public readonly lastName: string;

  public constructor(firstName: string, lastName: string) {
    this.firstName = firstName;
    this.lastName = lastName;
  }

  public get id(): string {
    if (!this._id) {
      throw new Error("id is empty");
    }
    return this._id;
  }
}
