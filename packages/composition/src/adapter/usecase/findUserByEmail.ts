import { UserRestClient } from "../../client/user";

interface FindUserByEmailCommandConstructor {
  email: string;
}

export class FindUserByEmailCommand {
  public readonly email: string;

  public constructor({ email }: FindUserByEmailCommandConstructor) {
    this.email = email;
  }
}

export class FindUserByEmailUseCase {
  public async execute(command: FindUserByEmailCommand) {
    const { email } = command;
    const userClient = new UserRestClient();
    return await userClient.findByEmail(email);
  }
}
