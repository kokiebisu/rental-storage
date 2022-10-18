import { RentalFeeType } from "../Domain/Enum";
import { Amount } from "../Domain/Model";
import { AmountInterface } from "./Amount";

export interface FeeConstructor {
  amount: Amount;
  type: RentalFeeType;
}

export interface FeeInterface {
  amount: AmountInterface;
  type: string;
}
