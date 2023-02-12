import { SpaceRestClient } from "../../client";
import { InternalServerError } from "../../error";

interface CreateSpaceCommandConstructor {
  lenderId: string;
  streetAddress: string;
  latitude: number;
  longitude: number;
  imageUrls: string[];
  title: string;
}

export class CreateSpaceCommand {
  public readonly lenderId: string;
  public readonly streetAddress: string;
  public readonly latitude: number;
  public readonly longitude: number;
  public readonly imageUrls: string[];
  public readonly title: string;

  public constructor({
    lenderId,
    streetAddress,
    latitude,
    longitude,
    imageUrls,
    title,
  }: CreateSpaceCommandConstructor) {
    this.lenderId = lenderId;
    this.streetAddress = streetAddress;
    this.latitude = latitude;
    this.longitude = longitude;
    this.imageUrls = imageUrls;
    this.title = title;
  }
}

export class CreateSpaceUseCase {
  public async execute(command: CreateSpaceCommand): Promise<{ uid: string }> {
    const { lenderId, streetAddress, latitude, longitude, imageUrls, title } =
      command;
    if (
      !lenderId ||
      !streetAddress ||
      !latitude ||
      !longitude ||
      !imageUrls ||
      !title
    ) {
      throw new InternalServerError();
    }
    const spaceClient = new SpaceRestClient();
    return await spaceClient.createSpace(
      lenderId,
      streetAddress,
      latitude,
      longitude,
      imageUrls,
      title
    );
  }
}
