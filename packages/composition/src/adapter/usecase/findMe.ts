import { UserRestClient } from "../../client";
import { InternalServerError } from "../../error";

interface FindMeCommandConstructor {
  id: string;
}

export class FindMeCommand {
  public readonly id: string;

  public constructor({ id }: FindMeCommandConstructor) {
    this.id = id;
  }
}

export class FindMeUseCase {
  public async execute(command: FindMeCommand): Promise<User> {
    const { id } = command;
    if (!id) {
      throw new InternalServerError();
    }
    const userClient = new UserRestClient();
    const data = await userClient.findUser(id);
    return {
      id: data.user.uid,
      firstName: data.user.firstName,
      lastName: data.user.lastName,
      emailAddress: data.user.emailAddress,
      streetAddress: data.user.streetAddress,
      createdAt: data.user.createdAt,
      updatedAt: data.user.updatedAt,
    };
  }
}
