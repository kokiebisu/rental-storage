import { SpaceRestClient } from "../../client";
import { InternalServerError } from "../../error";

interface CreateSpaceCommandConstructor {
  lenderId: string;
  imageUrls: string[];
  title: string;
  description: string;
  location: ILocation;
}

export class CreateSpaceCommand {
  public readonly lenderId: string;
  public readonly imageUrls: string[];
  public readonly title: string;
  public readonly description: string;
  public readonly location: ILocation;

  public constructor({
    lenderId,
    imageUrls,
    title,
    description,
    location,
  }: CreateSpaceCommandConstructor) {
    this.lenderId = lenderId;
    this.imageUrls = imageUrls;
    this.title = title;
    this.description = description;
    this.location = location;
  }
}

export class CreateSpaceUseCase {
  public async execute(command: CreateSpaceCommand): Promise<{ uid: string }> {
    const { lenderId, imageUrls, title, description, location } = command;
    if (!lenderId || !imageUrls || !title || !description || !location) {
      throw new InternalServerError();
    }
    const spaceClient = new SpaceRestClient();
    return await spaceClient.createSpace(
      lenderId,
      imageUrls,
      title,
      description,
      location
    );
  }
}
