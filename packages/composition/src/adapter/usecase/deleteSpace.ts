import { RestAPIClient } from "../../client";
import { InternalServerError } from "../../error";
import { SpaceResourceURLBuilder } from "../../resource";

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
    const client = new RestAPIClient();
    const builder = new SpaceResourceURLBuilder();
    const response = await client.delete<{ uid: string }>(
      builder.deleteSpace(id)
    );
    return { uid: response.data.uid };
  }
}
