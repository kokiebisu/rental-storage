import { UserRestClient } from "../../client";
import { InternalServerError } from "../../error";

interface RemoveUserByIdCommandConstructor {
  userId: string;
}

export class RemoveUserByIdCommand {
  public readonly userId: string;

  public constructor({ userId }: RemoveUserByIdCommandConstructor) {
    this.userId = userId;
  }
}

export class RemoveUserByIdUseCase {
  public async execute(
    command: RemoveUserByIdCommand
  ): Promise<{ uid: string } | undefined> {
    const { userId } = command;
    if (!userId) {
      throw new InternalServerError();
    }
    const userClient = new UserRestClient();
    // TODO: must remove all listings associated with the user
    const data = await userClient.removeUserById(userId);
    return { uid: data.uid };
  }
}
