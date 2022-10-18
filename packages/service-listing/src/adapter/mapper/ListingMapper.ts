import { Listing, StreetAddress } from "../../Domain/Model";
import { ListingInterface, ListingRawInterface } from "../../Types";

export class ListingMapper {
  public static toEntityFromRaw(data: ListingRawInterface): Listing {
    return new Listing({
      id: data.id,
      lenderId: data.lender_id,
      streetAddress: new StreetAddress(data.street_address),
      latitude: data.latitude,
      longitude: data.longitude,
      imageUrls: data.image_urls,
    });
  }

  public static toDTOFromRaw(data: ListingRawInterface): ListingInterface {
    return {
      id: data.id,
      uid: data.uid,
      lenderId: data.lender_id,
      streetAddress: data.street_address,
      latitude: data.latitude,
      longitude: data.longitude,
      imageUrls: data.image_urls,
    };
  }

  public static toDTOFromEntity(data: Listing): ListingInterface {
    return {
      ...(data.id && { id: data.id }),
      uid: data.uid,
      lenderId: data.lenderId,
      streetAddress: data.streetAddress.value,
      latitude: data.latitude,
      longitude: data.longitude,
      imageUrls: data.imageUrls,
    };
  }
}
