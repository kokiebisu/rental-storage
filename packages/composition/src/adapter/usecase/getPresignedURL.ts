import { ImageRestClient } from "../../client/image";

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
    const client = new ImageRestClient();
    return client.getPresignedURL(filename);
  }
}
