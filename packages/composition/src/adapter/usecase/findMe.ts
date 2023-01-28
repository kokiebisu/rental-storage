import { UserRestClient } from "../../client";
import { InternalServerError } from "../../error";

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
    if (!userId) {
      throw new InternalServerError();
    }
    const userClient = new UserRestClient();
    return await userClient.findUserById(userId);
  }
}
