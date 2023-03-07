import { RestAPIClient, SpaceResourceURLBuilder } from "../../client";

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
  public async execute(command: FindSpaceCommand): Promise<ISpace> {
    const { id } = command;
    const client = new RestAPIClient();
    const builder = new SpaceResourceURLBuilder();
    const response = await client.get<{
      space: Omit<ISpace, "id"> & { uid: string };
    }>(builder.findSpace(id));
    return {
      id: response.data.space.uid,
      ...response.data.space,
    };
  }
}
