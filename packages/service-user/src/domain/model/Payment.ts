import { PaymentConstructor } from "../types";

export enum ProviderType {
  STRIPE = "stripe",
}

export class Payment {
  private _id?: number;
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
    this._id = id;
    this.providerId = providerId;
    this.providerType = providerType;
    this.userId = userId;
  }

  public get id(): number {
    if (!this._id) {
      throw new Error("id doesn't exist in payment");
    }
    return this._id;
  }

  public set id(value: number) {
    this._id = value;
  }

  private isProviderType(value: string): value is ProviderType {
    return value === "stripe";
  }

  public static isPayment(data: any): data is Payment {
    return (
      (data as Payment).providerId !== undefined &&
      (data as Payment).userId !== undefined
    );
  }
}
