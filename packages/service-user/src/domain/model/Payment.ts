import { PaymentConstructor } from "../../types";

export enum ProviderType {
  STRIPE = "stripe",
}

export class Payment {
  public readonly id?: number;
  public readonly customerId: string;
  public readonly uid: string;
  public readonly providerType: ProviderType;
  public readonly userId: string;

  public constructor({
    id,
    uid,
    customerId,
    providerType,
    userId,
  }: PaymentConstructor) {
    if (!this.isProviderType(providerType)) {
      throw new Error("Provided providerType is not valid");
    }
    this.id = id;
    this.uid = uid;
    this.customerId = customerId;
    this.providerType = providerType;
    this.userId = userId;
  }

  private isProviderType(value: string): value is ProviderType {
    return value === "stripe";
  }
}
