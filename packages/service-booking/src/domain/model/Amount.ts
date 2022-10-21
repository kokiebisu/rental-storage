import { AmountConstructor } from "../../types";
import { Currency } from "../enums";

export class Amount {
  public readonly value: number;
  public readonly currency: string;

  public constructor({ value, currency }: AmountConstructor) {
    this.validateValue(value);
    this.validateCurrency(currency);
    this.value = value;
    this.currency = currency;
  }

  private validateValue(value: number) {
    if (value < 0) {
      throw new Error("Amount value needs to greater or equal to 0");
    }
  }

  private validateCurrency(value: any) {
    if (!Object.values(Currency).includes(value)) {
      throw new Error("Provided value is not a valid currency");
    }
  }
}
