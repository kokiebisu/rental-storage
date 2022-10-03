export interface StorageItemInterface {
  id?: number;
  name: string;
  imageUrls: string[];
  borrowerId: string;
  listingId: string;
  createdAt: string;
  updatedAt?: string;
}

export interface StorageItemRawInterface {
  id: number;
  uid: string;
  name: string;
  image_urls: string[];
  borrower_id: string;
  listing_id: string;
  created_at: string;
  updated_at?: string;
}

export interface StorageItemConstructor {
  id?: number;
  uid?: string;
  name: string;
  imageUrls?: string[];
  borrowerId: string;
  listingId: string;
  createdAt?: Date;
  updatedAt?: Date;
}
