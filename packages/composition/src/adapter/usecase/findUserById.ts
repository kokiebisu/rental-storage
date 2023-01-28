import { UserRestClient } from "../../client";
import { InternalServerError } from "../../error";

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
    if (!userId) {
      throw new InternalServerError();
    }
    const userClient = new UserRestClient();
    const response = await userClient.findUserById(userId);
    return response;
  }
}
