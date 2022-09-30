import { StorageItem, StorageItemRawInterface } from "../../../domain/model";
import { StorageItemInterface } from "../../../types";
import { TimeUtil } from "../../../utils";

export class StorageItemMapper {
  public static toDTOFromRaw(
    data: StorageItemRawInterface
  ): StorageItemInterface {
    return {
      id: data.id,
      name: data.name,
      imageUrls: data.image_urls,
      guestId: data.guest_id,
      listingId: data.listing_id,
      createdAt: data.created_at,
      ...(data.updated_at && { updatedAt: data.updated_at }),
    };
  }

  public static toDTOFromEntity(data: StorageItem): StorageItemInterface {
    return {
      ...(data.id && { id: data.id }),
      name: data.name,
      imageUrls: data.imageUrls,
      guestId: data.guestId,
      listingId: data.listingId,
      createdAt: TimeUtil.toDate(data.createdAt),
      ...(data.updatedAt && { updatedAt: TimeUtil.toDate(data.updatedAt) }),
    };
  }

  public static toDTOFromBookingStream(data: any): StorageItemInterface[] {
    const { items, listing_id, guest_id, created_at, updated_at } = data;

    return items.L.map((itemStringified: any) => {
      const item = JSON.parse(itemStringified.S);
      return {
        name: item.name,
        imageUrls: item.imageUrls.map((imageUrl: any) => imageUrl),
        guestId: guest_id.S,
        listingId: listing_id.S,
        createdAt: created_at.S,
        ...(updated_at && { updatedAt: updated_at.S }),
      };
    });
  }
}
