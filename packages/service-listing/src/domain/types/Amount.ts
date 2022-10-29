import { CurrencyType } from "../enum";

export interface AmountInterface {
  value: number;
  currency: string;
}

export interface AmountConstructor {
  value: number;
  currency: CurrencyType;
}
