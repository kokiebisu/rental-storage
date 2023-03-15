import { RestAPIClient } from "../../client";
import { SpaceResourceURLBuilder } from "../../resource";

interface FindSpacesCommandConstructor {
  userId?: string;
  offset?: number;
  limit?: number;
}

export class FindSpacesCommand {
  public readonly userId?: string;
  public readonly offset?: number;
  public readonly limit?: number;

  public constructor({ userId, offset, limit }: FindSpacesCommandConstructor) {
    this.userId = userId;
    this.offset = offset;
    this.limit = limit;
  }
}

export class FindSpacesUseCase {
  public async execute(command: FindSpacesCommand): Promise<ISpace[]> {
    const { userId, offset, limit } = command;
    const client = new RestAPIClient();
    const builder = new SpaceResourceURLBuilder();
    const param: {
      userId?: string;
      offset?: number;
      limit?: number;
    } = {};
    if (userId) {
      param["userId"] = userId;
    }
    if (offset !== undefined) {
      param["offset"] = offset;
    }
    if (limit !== undefined) {
      param["limit"] = limit;
    }
    const url = builder.findSpaces(param);

    const response = await client.get<{
      spaces: (Omit<ISpace, "id"> & { uid: string })[];
    }>(url);
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
