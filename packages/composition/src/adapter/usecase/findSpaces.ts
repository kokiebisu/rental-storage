import { SpaceRestClient } from "../../client";

interface FindSpacesCommandConstructor {
  userId: string;
}

export class FindSpacesCommand {
  public readonly userId: string;

  public constructor({ userId }: FindSpacesCommandConstructor) {
    this.userId = userId;
  }
}

export class FindSpacesUseCase {
  public async execute(command: FindSpacesCommand): Promise<Space[]> {
    const { userId } = command;
    const client = new SpaceRestClient();
    const data = await client.findSpaces(userId);
    return data.spaces.map((space: any) => {
      return {
        id: space.uid,
        ...space,
      };
    });
  }
}
