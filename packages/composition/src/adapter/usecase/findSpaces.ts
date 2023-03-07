import { RestAPIClient, SpaceResourceURLBuilder } from "../../client";

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
  public async execute(command: FindSpacesCommand): Promise<ISpace[]> {
    const { userId } = command;
    const client = new RestAPIClient();
    const response = await client.get<{
      spaces: (Omit<ISpace, "id"> & { uid: string })[];
    }>(SpaceResourceURLBuilder.findSpaces({ userId }));
    return response.data.spaces.map(
      (space: Omit<ISpace, "id"> & { uid: string }) => {
        return {
          id: space.uid,
          ...space,
        };
      }
    );
  }
}
