import { RentalFeeType } from "../enum";
import { Amount } from "../model";
import { AmountInterface } from "./Amount";

export interface FeeConstructor {
  amount: Amount;
  type: RentalFeeType;
}

export interface FeeInterface {
  amount: AmountInterface;
  type: string;
}
