export interface ItemInterface {
  id?: number;
  name: string;
  imageUrls: string[];
  ownerId: string;
  listingId: string;
  createdAt: string;
  updatedAt?: string;
}

export interface ItemRawInterface {
  id: number;
  uid: string;
  name: string;
  image_urls: string[];
  owner_id: string;
  listing_id: string;
  created_at: string;
  updated_at?: string;
}

export interface ItemConstructor {
  id?: number;
  uid?: string;
  name: string;
  imageUrls?: string[];
  ownerId: string;
  listingId: string;
  createdAt?: Date;
  updatedAt?: Date;
}
