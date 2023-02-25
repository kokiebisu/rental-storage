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
  public async execute(command: FindSpaceCommand): Promise<ISpace> {
    const { id } = command;
    const client = new SpaceRestClient();
    const data = await client.findSpace(id);
    return {
      id: data.space.uid,
      ...data.space,
    };
  }
}
