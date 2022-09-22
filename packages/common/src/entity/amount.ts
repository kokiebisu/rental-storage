import { Currency } from "../enum";

export interface Amount {
  value: number;
  currency: Currency;
}
