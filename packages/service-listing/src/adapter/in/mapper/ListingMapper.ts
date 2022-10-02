import { Listing } from "../../../domain/model";
import {
  AggregatedListingInterface,
  HostInterface,
  ListingInterface,
  ListingRawInterface,
  StorageItemInterface,
} from "../../../types";

export class ListingMapper {
  public static toAggregatedDTOFromRaw(
    data: ListingRawInterface
  ): AggregatedListingInterface {
    // aggregate host data by the host data
    const host: HostInterface = {
      id: data.host_id,
      firstName: "mock host firstName",
      lastName: "mock host lastName",
    };
    const items: StorageItemInterface[] = [
      {
        name: "random item name",
        imageUrls: ["random_imageUrls"],
        userId: "1",
        listingId: "1",
      },
    ]; // must fetch from api endpoint
    return {
      id: data.id,
      host,
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
      hostId: data.host_id,
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
    // aggregate host data by the host data
    const host: HostInterface = {
      id: data.hostId,
      firstName: "mock host firstName",
      lastName: "mock host lastName",
    };

    const items: StorageItemInterface[] = [
      {
        name: "random item name",
        imageUrls: ["random_imageUrls"],
        userId: "1",
        listingId: "1",
      },
    ];

    return {
      id: data.id,
      host,
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
      hostId: data.hostId,
      streetAddress: data.streetAddress.value,
      latitude: data.latitude,
      longitude: data.longitude,
      imageUrls: data.imageUrls,
      items: data.items,
    };
  }
}
