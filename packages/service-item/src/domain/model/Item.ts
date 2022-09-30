export interface StorageItemRawInterface {
  id: string;
  name: string;
  image_urls: string[];
  guest_id: string;
  listing_id: string;
  created_at: string;
  updated_at?: string;
}

interface StorageItemConstructor {
  id?: string;
  name: string;
  imageUrls?: string[];
  guestId: string;
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
  public readonly guestId: string;
  public readonly listingId: string;

  public constructor({
    name,
    imageUrls = [],
    createdAt = new Date(),
    updatedAt,
    guestId,
    listingId,
  }: StorageItemConstructor) {
    this.name = name;
    this.imageUrls = imageUrls;
    this.createdAt = createdAt;
    this.updatedAt = updatedAt;
    this.guestId = guestId;
    this.listingId = listingId;
  }
}
