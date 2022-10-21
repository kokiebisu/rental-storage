import { RentalFeeType } from "../domain/enum";
import { Amount } from "../domain/model";
import { AmountInterface } from "./Amount";

export interface FeeConstructor {
  amount: Amount;
  type: RentalFeeType;
}

export interface FeeInterface {
  amount: AmountInterface;
  type: string;
}
