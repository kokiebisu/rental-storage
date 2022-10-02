import { v4 as uuid } from "uuid";
import { StorageItemConstructor } from "../../types";

export class StorageItem {
  public readonly id?: number;
  public readonly uid?: string;
  public readonly name: string;
  public readonly imageUrls: string[];
  public readonly createdAt: Date;
  public readonly updatedAt?: Date | undefined;
  public readonly guestId: string;
  public readonly listingId: string;

  public constructor({
    id,
    name = uuid(),
    imageUrls = [],
    createdAt = new Date(),
    updatedAt,
    guestId,
    listingId,
  }: StorageItemConstructor) {
    this.id = id;
    this.name = name;
    this.imageUrls = imageUrls;
    this.createdAt = createdAt;
    this.updatedAt = updatedAt;
    this.guestId = guestId;
    this.listingId = listingId;
  }
}
