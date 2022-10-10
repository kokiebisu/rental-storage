import { Listing } from "../../../domain/model";
import {
  AggregatedListingInterface,
  LenderInterface,
  ListingInterface,
  ListingRawInterface,
  ItemInterface,
} from "../../../types";

export class ListingMapper {
  public static toAggregatedDTOFromRaw(
    data: ListingRawInterface
  ): AggregatedListingInterface {
    // aggregate lender data by the lender data
    const lender: LenderInterface = {
      id: data.lender_id,
      firstName: "mock lender firstName",
      lastName: "mock lender lastName",
    };
    const items: ItemInterface[] = [
      {
        name: "random item name",
        imageUrls: ["random_imageUrls"],
        userId: "1",
        listingId: "1",
      },
    ]; // must fetch from api endpoint
    return {
      id: data.id,
      lender,
      streetAddress: data.street_address,
      latitude: data.latitude,
      longitude: data.longitude,
      imageUrls: data.image_urls,
      items,
    };
  }

  public static toDTOFromRaw(data: ListingRawInterface): ListingInterface {
    return {
      id: data.id,
      lenderId: data.lender_id,
      streetAddress: data.street_address,
      latitude: data.latitude,
      longitude: data.longitude,
      imageUrls: data.image_urls,
      items: data.items,
    };
  }

  public static toAggregated(
    data: ListingInterface
  ): AggregatedListingInterface {
    // aggregate lender data by the lender data
    const lender: LenderInterface = {
      id: data.lenderId,
      firstName: "mock lender firstName",
      lastName: "mock lender lastName",
    };

    const items: ItemInterface[] = [
      {
        name: "random item name",
        imageUrls: ["random_imageUrls"],
        userId: "1",
        listingId: "1",
      },
    ];

    return {
      id: data.id,
      lender,
      streetAddress: data.streetAddress,
      latitude: data.latitude,
      longitude: data.longitude,
      imageUrls: data.imageUrls,
      items,
    };
  }

  public static toDTOFromEntity(data: Listing): ListingInterface {
    return {
      ...(data.id && { id: data.id }),
      lenderId: data.lenderId,
      streetAddress: data.streetAddress.value,
      latitude: data.latitude,
      longitude: data.longitude,
      imageUrls: data.imageUrls,
      items: data.items,
    };
  }
}
