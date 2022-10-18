import { ItemConstructor } from "../../Types";

export class Item {
  private _id?: number;
  public readonly uid?: string;
  public readonly name: string;
  public readonly imageUrls: string[];
  public readonly createdAt: Date;
  public readonly updatedAt?: Date | undefined;
  public readonly ownerId: string;
  public readonly listingId: string;

  public constructor({
    id,
    name,
    imageUrls = [],
    createdAt = new Date(),
    updatedAt,
    ownerId,
    listingId,
  }: ItemConstructor) {
    this._id = id;
    this.name = name;
    this.imageUrls = imageUrls;
    this.createdAt = createdAt;
    this.updatedAt = updatedAt;
    this.ownerId = ownerId;
    this.listingId = listingId;
  }

  public get id() {
    if (!this._id) {
      throw new Error("id is not defined");
    }
    return this._id;
  }

  public set id(value: number) {
    this._id = value;
  }

  public static isItem(data: any): data is Item {
    return (
      (data as Item).ownerId !== undefined &&
      (data as Item).listingId !== undefined
    );
  }
}
