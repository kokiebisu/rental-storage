import { UserRestClient } from "../../client";
import { InternalServerError } from "../../error";

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
    const userClient = new UserRestClient();
    // TODO: must remove all listings associated with the user
    const data = await userClient.deleteUser(id);
    return { id: data.uid };
  }
}
