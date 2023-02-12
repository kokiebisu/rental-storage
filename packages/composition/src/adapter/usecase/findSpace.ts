import { SpaceRestClient } from "../../client";

interface FindSpaceCommandConstructor {
  id: string;
}

export class FindSpaceCommand {
  public readonly id: string;

  public constructor({ id }: FindSpaceCommandConstructor) {
    this.id = id;
  }
}

export class FindSpaceUseCase {
  public async execute(command: FindSpaceCommand): Promise<Space> {
    const { id } = command;
    const client = new SpaceRestClient();
    const data = await client.findSpace(id);
    return {
      id: data.space.uid,
      feeType: data.space.feeType,
      feeAmount: data.space.feeAmount,
      feeCurrency: data.space.feeCurrency,
      imageUrls: data.space.imageUrls,
      latitude: data.space.latitude,
      longitude: data.space.longitude,
      lenderId: data.space.lenderId,
      streetAddress: data.space.streetAddress,
      title: data.space.title,
    };
  }
}
