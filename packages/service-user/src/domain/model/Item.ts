import { ItemConstructor } from "../../types";

export class Item {
  public readonly id?: number;
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
    this.id = id;
    this.name = name;
    this.imageUrls = imageUrls;
    this.createdAt = createdAt;
    this.updatedAt = updatedAt;
    this.ownerId = ownerId;
    this.listingId = listingId;
  }
}
