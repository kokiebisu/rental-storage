import { BookingInterface } from "@rental-storage-project/common";
import { Booking } from "../../../domain/model";

export class BookingMapper {
  public static toDTOFromEntity(entity: Booking): BookingInterface {
    return {
      id: entity.id,
      status: entity.status.toString(),
      amount: entity.amount,
      userId: entity.userId,
      listingId: entity.listingId,
      createdAt: entity.createdAt.toLocaleString(),
      ...(entity.updatedAt && { updatedAt: entity.updatedAt.toLocaleString() }),
      items: entity.items,
    };
  }

  public static toDTOFromRaw(raw: any): BookingInterface {
    const {
      Item: {
        id,
        status,
        listing_id,
        user_id,
        created_at,
        items,
        amount,
        updated_at,
      },
    } = raw;

    const itemsParsed = items.L.map((item: any) => JSON.parse(item));

    return {
      id: id.S,
      status: status.S,
      amount: {
        value: parseInt(amount.M.value.N),
        currency: amount.M.currency.S,
      },
      userId: user_id.S,
      listingId: listing_id.S,
      createdAt: created_at.S,
      ...(updated_at && { updatedAt: updated_at.S }),
      items: itemsParsed,
    };
  }
}
