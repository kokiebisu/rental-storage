export interface StorageItemRawInterface {
  id: string;
  name: string;
  image_urls: string[];
  user_id: string;
  listing_id: string;
  created_at: string;
  updated_at?: string;
}

interface StorageItemConstructor {
  id?: string;
  name: string;
  imageUrls?: string[];
  userId: string;
  listingId: string;
  createdAt?: Date;
  updatedAt?: Date;
}

export class StorageItem {
  public readonly id?: string;
  public readonly name: string;
  public readonly imageUrls: string[];
  public readonly createdAt: Date;
  public readonly updatedAt?: Date | undefined;
  public readonly userId: string;
  public readonly listingId: string;

  public constructor({
    name,
    imageUrls = [],
    createdAt = new Date(),
    updatedAt,
    userId,
    listingId,
  }: StorageItemConstructor) {
    this.name = name;
    this.imageUrls = imageUrls;
    this.createdAt = createdAt;
    this.updatedAt = updatedAt;
    this.userId = userId;
    this.listingId = listingId;
  }
}
