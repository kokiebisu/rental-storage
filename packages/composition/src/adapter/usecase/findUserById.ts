import { UserRestClient } from "../../client/user";

interface FindUserByIdCommandConstructor {
  userId: string;
}

export class FindUserByIdCommand {
  public readonly userId: string;

  public constructor({ userId }: FindUserByIdCommandConstructor) {
    this.userId = userId;
  }
}

export class FindUserByIdUseCase {
  public async execute(command: FindUserByIdCommand) {
    const { userId } = command;
    const userClient = new UserRestClient();
    return await userClient.findUserById(userId);
  }
}
