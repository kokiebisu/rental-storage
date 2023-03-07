import {
  RestAPIClient,
  UserResourceURLBuilder,
  UserRestClient,
} from "../../client";
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
    const client = new RestAPIClient();
    // TODO: must remove all spaces associated with the user
    const response = await client.delete<{ uid: string }>(
      UserResourceURLBuilder.deleteUser(id)
    );
    return { id: response.data.uid };
  }
}
