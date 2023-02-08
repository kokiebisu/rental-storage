import { ImageRestClient } from "../../client";
import { InternalServerError } from "../../error";

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
    const client = new ImageRestClient();
    return await client.getPresignedURL(filename);
  }
}
