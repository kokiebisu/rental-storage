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

export interface GuestInterface {
  id?: string;
  firstName: string;
  lastName: string;
}

export class Guest extends User {
  public constructor(firstName: string, lastName: string) {
    super(firstName, lastName);
  }

  public toDTO(): GuestInterface {
    return {
      firstName: this.firstName,
      lastName: this.lastName,
    };
  }
}

export interface HostInterface {
  id?: string;
  firstName: string;
  lastName: string;
}

export class Host extends User {
  public constructor(firstName: string, lastName: string) {
    super(firstName, lastName);
  }

  public toDTO(): HostInterface {
    return {
      firstName: this.firstName,
      lastName: this.lastName,
    };
  }
}
