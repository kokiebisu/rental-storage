import { UserRestClient } from "../../client/user";

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
  public async execute(command: RemoveUserByIdCommand) {
    const { userId } = command;
    const userClient = new UserRestClient();
    return await userClient.removeUserById(userId);
  }
}
