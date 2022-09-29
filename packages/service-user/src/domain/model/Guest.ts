import { User } from "./User";

export interface GuestRawInterface {
  id: string;
  first_name: string;
  last_name: string;
  items: string[];
}

interface GuestConstructor {
  firstName: string;
  lastName: string;
  items?: string[];
}

export class Guest extends User {
  private _items: string[];

  public constructor({ firstName, lastName, items = [] }: GuestConstructor) {
    super(firstName, lastName);
    this._items = items;
  }

  public get items(): string[] {
    return this._items;
  }
  public set items(value: string[]) {
    this._items = value;
  }
}
