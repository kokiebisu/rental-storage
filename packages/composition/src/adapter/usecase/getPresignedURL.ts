import { RestAPIClient } from "../../client";
import { InternalServerError } from "../../error";
import { ImageResourceURLBuilder } from "../../resource";

interface GetPresignedURLCommandConstructor {
  filename: string;
}

export class GetPresignedURLCommand {
  public readonly filename: string;

  public constructor({ filename }: GetPresignedURLCommandConstructor) {
    this.filename = filename;
  }
}

export class GetPresignedURLUseCase {
  public async execute(command: GetPresignedURLCommand) {
    const { filename } = command;
    if (!filename) {
      throw new InternalServerError();
    }
    const client = new RestAPIClient();
    const builder = new ImageResourceURLBuilder();
    return await client.get(builder.getPresignedURL(filename));
  }
}
