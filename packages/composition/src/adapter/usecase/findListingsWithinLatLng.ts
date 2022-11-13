import { ListingRestClient } from "../../client/listing";

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
    const listingClient = new ListingRestClient();
    const response = await listingClient.findListingsWithinLatLng(
      latitude,
      longitude,
      range
    );
    return response.data;
  }
}