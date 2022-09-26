export interface StorageItemInterface {
  id?: string;
  name: string;
  imageUrls: string[];
  userId: string;
  listingId: string;
  createdAt?: Date;
  updatedAt?: Date;
}
