export interface StorageItemInterface {
  id?: number;
  name: string;
  imageUrls: string[];
  userId: number;
  listingId: number;
  createdAt?: Date;
  updatedAt?: Date;
}
