import { FeeConstructor } from "../../Types";
import { RentalFeeType } from "../Enum";
import { Amount } from "./Amount";

export class Fee {
  public readonly amount: Amount;
  public readonly type: RentalFeeType;

  public constructor({ amount, type }: FeeConstructor) {
    this.validateType(type);
    this.amount = amount;
    this.type = type;
  }

  private validateType(value: RentalFeeType) {
    if (!Object.values(RentalFeeType).includes(value)) {
      throw new Error("Provided type is not a member of the RentalFeeType");
    }
  }
}
