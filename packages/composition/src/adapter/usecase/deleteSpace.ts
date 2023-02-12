import { SpaceRestClient } from "../../client";
import { InternalServerError } from "../../error";

interface DeleteSpaceCommandConstructor {
  id: string;
}

export class DeleteSpaceCommand {
  public readonly id: string;

  public constructor({ id }: DeleteSpaceCommandConstructor) {
    this.id = id;
  }
}

export class DeleteSpaceUseCase {
  public async execute(command: DeleteSpaceCommand): Promise<{ uid: string }> {
    const { id } = command;
    if (!id) {
      throw new InternalServerError();
    }
    const spaceClient = new SpaceRestClient();
    return await spaceClient.deleteSpace(id);
  }
}
