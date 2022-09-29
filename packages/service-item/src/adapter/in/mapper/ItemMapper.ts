import { StorageItem, StorageItemRawInterface } from "../../../domain/model";
import { StorageItemInterface } from "../../../types";

export class StorageItemMapper {
  public static toDTOFromRaw(data: StorageItemRawInterface) {
    return {
      id: data.id,
      name: data.name,
      imageUrls: data.image_urls,
      userId: data.user_id,
      listingId: data.listing_id,
      createdAt: new Date(data.created_at),
      ...(data.updated_at && { updatedAt: new Date(data.updated_at) }),
    };
  }

  public static toDTOFromEntity(data: StorageItem): StorageItemInterface {
    return {
      ...(data.id && { id: data.id }),
      name: data.name,
      imageUrls: data.imageUrls,
      userId: data.userId,
      listingId: data.listingId,
    };
  }
}
