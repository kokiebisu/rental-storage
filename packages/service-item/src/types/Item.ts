export interface StorageItemInterface {
  id?: string;
  name: string;
  imageUrls: string[];
  guestId: string;
  listingId: string;
  createdAt: string;
  updatedAt?: string;
}
