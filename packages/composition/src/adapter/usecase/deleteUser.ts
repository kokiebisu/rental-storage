import { RestAPIClient } from "../../client";
import { InternalServerError } from "../../error";
import { UserResourceURLBuilder } from "../../resource";

interface DeleteUserCommandConstructor {
  id: string;
}

export class DeleteUserCommand {
  public readonly id: string;

  public constructor({ id }: DeleteUserCommandConstructor) {
    this.id = id;
  }
}

export class DeleteUserUseCase {
  public async execute(command: DeleteUserCommand): Promise<{ id: string }> {
    const { id } = command;
    if (!id) {
      throw new InternalServerError();
    }
    const client = new RestAPIClient();
    const builder = new UserResourceURLBuilder();
    // TODO: must remove all spaces associated with the user
    const response = await client.delete<{ uid: string }>(
      builder.deleteUser(id)
    );
    return { id: response.data.uid };
  }
}
