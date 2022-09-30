import { Host } from "../../../domain/model";
import { HostInterface, HostRawInterface } from "../../../types";
import { TimeUtil } from "../../../utils";

export class HostMapper {
  public static toDTOFromRaw(data: HostRawInterface) {
    return {
      id: data.id,
      firstName: data.first_name,
      lastName: data.last_name,
      createdAt: data.created_at,
      ...(data.updated_at && { updatedAt: data.updated_at }),
    };
  }

  public static toDTOFromEntity(data: Host): HostInterface {
    return {
      ...(data.id && { id: data.id }),
      firstName: data.firstName,
      lastName: data.lastName,
      createdAt: TimeUtil.toDate(data.createdAt),
      ...(data.updatedAt && { updatedAt: TimeUtil.toDate(data.updatedAt) }),
    };
  }
}
