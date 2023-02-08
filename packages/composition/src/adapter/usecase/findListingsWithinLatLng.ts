import { ListingRestClient } from "../../client";
import { InternalServerError } from "../../error";

interface FindListingsWithinLatLngCommandConstructor {
  latitude: number;
  longitude: number;
  range: number;
}

export class FindListingsWithinLatLngCommand {
  public readonly latitude: number;
  public readonly longitude: number;
  public readonly range: number;

  public constructor({
    latitude,
    longitude,
    range,
  }: FindListingsWithinLatLngCommandConstructor) {
    this.latitude = latitude;
    this.longitude = longitude;
    this.range = range;
  }
}

export class FindListingsWithinLatLngUseCase {
  public async execute(command: FindListingsWithinLatLngCommand) {
    const { latitude, longitude, range } = command;
    if (!latitude || !longitude || !range) {
      throw new InternalServerError();
    }
    const listingClient = new ListingRestClient();
    const data = await listingClient.findListingsWithinLatLng(
      latitude,
      longitude,
      range
    );
    return data?.listings;
  }
}
