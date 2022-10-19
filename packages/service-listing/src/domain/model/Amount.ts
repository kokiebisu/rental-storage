import { AmountConstructor } from "../../types";
import { CurrencyType } from "../enum";

export class Amount {
  public readonly value: number;
  public readonly currency: CurrencyType;

  public constructor({ value, currency }: AmountConstructor) {
    this.validateValue(value);
    this.validateCurrency(currency);
    this.value = value;
    this.currency = currency;
  }

  private validateValue(value: number) {
    if (!value) {
      throw new Error("Amount value is undefined");
    }
  }

  private validateCurrency(value: CurrencyType) {
    if (!Object.values(CurrencyType).includes(value)) {
      throw new Error("Value doesn't belong to any currency types");
    }
  }
}
