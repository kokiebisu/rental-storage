import { UserRestClient } from "../../client";
import { InternalServerError } from "../../error";

interface FindUserByEmailCommandConstructor {
  emailAddress: string;
}

export class FindUserByEmailCommand {
  public readonly emailAddress: string;

  public constructor({ emailAddress }: FindUserByEmailCommandConstructor) {
    this.emailAddress = emailAddress;
  }
}

export class FindUserByEmailUseCase {
  public async execute(command: FindUserByEmailCommand) {
    const { emailAddress } = command;
    if (!emailAddress) {
      throw new InternalServerError();
    }
    const userClient = new UserRestClient();
    const data = await userClient.findByEmail(emailAddress);
    return data?.user;
  }
}
