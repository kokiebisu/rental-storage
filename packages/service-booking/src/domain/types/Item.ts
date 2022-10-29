export interface ItemInterface {
  id?: number;
  name: string;
  imageUrls: string[];
  userId: string;
  listingId: string;
  createdAt: Date;
  updatedAt?: Date;
}
