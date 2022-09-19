import { HostInterface } from "../../../user/src/entity/host";
import {
  AggregatedListingInterface,
  Listing,
  ListingInterface,
  ListingRawInterface,
} from "../entity";

export class ListingMapper {
  public static toAggregatedDTOFromRaw(
    data: ListingRawInterface
  ): ListingInterface {
    // aggregate host data by the host data
    const host: HostInterface = data.host_id;
    return {
      ...(data.id && { id: data.id }),
      host,
      emailAddress: data.email_address,
      streetAddress: data.street_address,
      latitude: data.latitude,
      longitude: data.longitude,
    };
  }

  public static toAggregatedDTO(
    data: ListingInterface
  ): AggregatedListingInterface {
    // aggregate host data by the host data
    const host = data.hostId;

    return {
      ...(data.id && { id: data.id }),
      host,
      emailAddress: data.emailAddress,
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
      emailAddress: data.emailAddress,
      latitude: data.latitude,
      longitude: data.longitude,
    };
  }
}
