import { Booking } from "../../domain/model";
import { BookingInterface } from "../../types";

export class BookingMapper {
  public static toDTOFromEntity(entity: Booking): BookingInterface {
    return {
      id: entity.id,
      status: entity.status.toString(),
      amount: entity.amount,
      ownerId: entity.ownerId,
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
        owner_id,
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
      ownerId: owner_id.S,
      listingId: listing_id.S,
      createdAt: created_at.S,
      ...(updated_at && { updatedAt: updated_at.S }),
      items: itemsParsed,
    };
  }
}
