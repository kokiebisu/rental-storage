import {
  AggregatedListingInterface,
  HostInterface,
  ListingInterface,
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
    return {
      id: data.id,
      host,
      streetAddress: data.street_address,
      latitude: data.latitude,
      longitude: data.longitude,
    };
  }

  public static toDTOFromRaw(data: ListingRawInterface): ListingInterface {
    return {
      id: data.id,
      hostId: data.host_id,
      streetAddress: data.street_address,
      latitude: data.latitude,
      longitude: data.longitude,
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

    return {
      id: data.id,
      host,
      streetAddress: data.streetAddress,
      latitude: data.latitude,
      longitude: data.longitude,
    };
  }

  public static toDTOFromEntity(data: Listing): ListingInterface {
    return {
      ...(data.id && { id: data.id }),
      hostId: data.hostId,
      streetAddress: data.streetAddress,
      latitude: data.latitude,
      longitude: data.longitude,
    };
  }
}
