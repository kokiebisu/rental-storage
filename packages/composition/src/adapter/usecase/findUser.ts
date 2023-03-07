import { RestAPIClient } from "../../client";
import UserResourceURLBuilder from "../../client/user";
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
  public async execute(command: FindUserCommand): Promise<IUser> {
    const { id } = command;
    if (!id) {
      throw new InternalServerError();
    }
    const client = new RestAPIClient();
    const builder = new UserResourceURLBuilder();
    const response = await client.get<{
      user: Omit<IUser, "id"> & { uid: string };
    }>(builder.findUser(id));
    return {
      ...response.data.user,
      id: response.data.user.uid,
    };
  }
}
