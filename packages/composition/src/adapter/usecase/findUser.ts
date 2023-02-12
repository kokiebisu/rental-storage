import { UserRestClient } from "../../client";
import { InternalServerError } from "../../error";

interface FindUserCommandConstructor {
  id: string;
}

export class FindUserCommand {
  public readonly id: string;

  public constructor({ id }: FindUserCommandConstructor) {
    this.id = id;
  }
}

export class FindUserUseCase {
  public async execute(command: FindUserCommand) {
    const { id } = command;
    if (!id) {
      throw new InternalServerError();
    }
    const userClient = new UserRestClient();
    const data = await userClient.findUser(id);
    return data?.user;
  }
}
