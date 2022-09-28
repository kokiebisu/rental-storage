import {
  AggregatedListingInterface,
  HostInterface,
  ListingInterface,
  StorageItemInterface,
} from "@rental-storage-project/common";
import { Listing, ListingRawInterface } from "../entity";

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
      items,
    };
  }

  public static toDTOFromEntity(data: Listing): ListingInterface {
    return {
      ...(data.id && { id: data.id }),
      hostId: data.hostId,
      streetAddress: data.streetAddress,
      latitude: data.latitude,
      longitude: data.longitude,
      items: data.items,
    };
  }
}
