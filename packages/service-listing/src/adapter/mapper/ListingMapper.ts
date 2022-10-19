import { CurrencyType, RentalFeeType } from "../../domain/enum";
import { Amount, Fee, Listing, StreetAddress } from "../../domain/model";
import { ListingInterface, ListingRawInterface } from "../../types";

export class ListingMapper {
  public static toEntityFromRaw(data: ListingRawInterface): Listing {
    return new Listing({
      id: data.id,
      lenderId: data.lender_id,
      title: data.title,
      streetAddress: new StreetAddress(data.street_address),
      latitude: data.latitude,
      longitude: data.longitude,
      imageUrls: data.image_urls,
      fee: new Fee({
        amount: new Amount({
          value: data.fee_amount,
          currency: data.fee_currency as CurrencyType,
        }),
        type: data.fee_type as RentalFeeType,
      }),
    });
  }

  public static toDTOFromEntity(data: Listing): ListingInterface {
    return {
      ...(data.id && { id: data.id }),
      uid: data.uid,
      title: data.title,
      lenderId: data.lenderId,
      streetAddress: data.streetAddress.value,
      latitude: data.latitude,
      longitude: data.longitude,
      imageUrls: data.imageUrls,
      fee: {
        amount: {
          value: data.fee.amount.value,
          currency: data.fee.amount.currency,
        },
        type: data.fee.type,
      },
    };
  }
}
