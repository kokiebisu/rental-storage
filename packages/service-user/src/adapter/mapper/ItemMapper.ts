import { Item } from "../../domain/model";
import { ItemInterface, ItemRawInterface } from "../../domain/types";
import { TimeUtil } from "../../utils";

export class ItemMapper {
  public static toEntityFromRaw(data: ItemRawInterface): Item {
    return new Item({
      id: data.id,
      name: data.name,
      imageUrls: data.image_urls,
      createdAt: new Date(data.created_at),
      updatedAt: new Date(data.created_at),
      ownerId: data.owner_id,
      listingId: data.listing_id,
    });
  }

  public static toDTOFromRaw(data: ItemRawInterface): ItemInterface {
    return {
      id: data.id,
      name: data.name,
      imageUrls: data.image_urls,
      ownerId: data.owner_id,
      listingId: data.listing_id,
      createdAt: data.created_at,
      ...(data.updated_at && { updatedAt: data.updated_at }),
    };
  }

  public static toDTOFromEntity(data: Item): ItemInterface {
    return {
      ...(data.id && { id: data.id }),
      name: data.name,
      imageUrls: data.imageUrls,
      ownerId: data.ownerId,
      listingId: data.listingId,
      createdAt: TimeUtil.toDate(data.createdAt),
      ...(data.updatedAt && { updatedAt: TimeUtil.toDate(data.updatedAt) }),
    };
  }

  public static toDTOFromBookingStream(data: any): ItemInterface[] {
    const { items, listing_id, owner_id, created_at, updated_at } = data;

    return items.L.map((itemStringified: any) => {
      const item = JSON.parse(itemStringified.S);
      return {
        name: item.name,
        imageUrls: item.imageUrls.map((imageUrl: any) => imageUrl),
        ownerId: Number(owner_id.N),
        listingId: Number(listing_id.N),
        createdAt: created_at.S,
        ...(updated_at && { updatedAt: updated_at.S }),
      };
    });
  }
}
