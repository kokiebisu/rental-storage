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
  private _id?: string;
  private _name: string;
  private _imageUrls: string[];
  private _createdAt?: Date | undefined;
  private _updatedAt?: Date | undefined;
  private _userId: string;
  private _listingId: string;

  public constructor({
    name,
    imageUrls = [],
    createdAt = new Date(),
    updatedAt,
    userId,
    listingId,
  }: StorageItemConstructor) {
    this._name = name;
    this._imageUrls = imageUrls;
    this._createdAt = createdAt;
    this._updatedAt = updatedAt;
    this._userId = userId;
    this._listingId = listingId;
  }

  public get id(): string | undefined {
    return this._id;
  }
  public set id(value: string | undefined) {
    this._id = value;
  }

  public get name(): string {
    return this._name;
  }
  public set name(value: string) {
    this._name = value;
  }

  public get imageUrls(): string[] {
    return this._imageUrls;
  }
  public set imageUrls(value: string[]) {
    this._imageUrls = value;
  }

  public get createdAt(): Date | undefined {
    return this._createdAt;
  }
  public set createdAt(value: Date | undefined) {
    this._createdAt = value;
  }

  public get updatedAt(): Date | undefined {
    return this._updatedAt;
  }
  public set updatedAt(value: Date | undefined) {
    this._updatedAt = value;
  }

  public get userId(): string {
    return this._userId;
  }
  public set userId(value: string) {
    this._userId = value;
  }

  public get listingId(): string {
    return this._listingId;
  }
  public set listingId(value: string) {
    this._listingId = value;
  }
}
