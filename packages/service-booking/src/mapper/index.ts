import { BookingInterface } from "@rental-storage-project/common";
import { Booking } from "../entity";

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
    };
  }

  public static toDTOFromRaw(raw: any): BookingInterface {
    const {
      Item: { id, status, listing_id, user_id, created_at, amount, updated_at },
    } = raw;
    // {
    //   Item: {
    //     listing_id: { S: '1' },
    //     user_id: { S: '1' },
    //     created_at: { S: '9/21/2022, 9:22:18 PM' },
    //     status: { S: '0' },
    //     amount: { M: [Object] },
    //     id: { S: '8f712a7d-f3ee-4761-90b2-e663de4623a5' }
    //   }
    // }
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
    };
  }
}
