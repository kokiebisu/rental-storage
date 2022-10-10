import { PaymentConstructor } from "../../types";

export enum ProviderType {
  STRIPE = "stripe",
}

export class Payment {
  public readonly id?: number;
  public readonly providerId: string;
  public readonly providerType: ProviderType;
  public readonly userId: number;

  public constructor({
    id,
    providerId,
    providerType,
    userId,
  }: PaymentConstructor) {
    if (!this.isProviderType(providerType)) {
      throw new Error("Provided providerType is not valid");
    }
    this.id = id;
    this.providerId = providerId;
    this.providerType = providerType;
    this.userId = userId;
  }

  private isProviderType(value: string): value is ProviderType {
    return value === "stripe";
  }
}
