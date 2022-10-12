export enum NameType {
  FirstName = "First",
  LastName = "Last",
}

export class Name {
  private _value: string;
  public readonly nameType: NameType;

  public constructor(nameType: NameType, value: string) {
    this.validateName(value, nameType);
    this._value = value;
    this.nameType = nameType;
  }

  private validateName(value: string, nameType: NameType) {
    if (!value) {
      throw new Error(`${nameType} name is not provided`);
    }
    if (value.length == 0) {
      throw new Error(`${nameType} name cannot be empty`);
    }
    // if (new RegExp(/^[a-z ,.'-]+$/i).test(value)) {
    //   throw new Error("Provided first name is invalid");
    // }
  }

  public get value() {
    return this.capitalize(this._value);
  }

  private capitalize(value: string) {
    return value[0].toUpperCase() + value.slice(1).toLowerCase();
  }
}
