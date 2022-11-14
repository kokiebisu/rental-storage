import { UserRestClient } from "../../client/user";

interface FindMeCommandConstructor {
  userId: string;
}

export class FindMeCommand {
  public readonly userId: string;

  public constructor({ userId }: FindMeCommandConstructor) {
    this.userId = userId;
  }
}

export class FindMeUseCase {
  public async execute(command: FindMeCommand) {
    const { userId } = command;
    const userClient = new UserRestClient();
    return await userClient.findUserById(userId);
  }
}
