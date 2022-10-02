import { Host } from "../../../domain/model";
import { HostInterface, HostRawInterface } from "../../../types";
import { TimeUtil } from "../../../utils";

export class HostMapper {
  public static toDTOFromRaw(data: HostRawInterface): HostInterface {
    return {
      ...(data.id && { id: data.id }),
      ...(data.uid && { uid: data.uid }),
      emailAddress: data.email_address,
      password: data.password,
      firstName: data.first_name,
      lastName: data.last_name,
      createdAt: data.created_at,
      ...(data.updated_at && { updatedAt: data.updated_at }),
    };
  }

  public static toDTOFromEntity(data: Host): HostInterface {
    return {
      id: data.id,
      uid: data.uid,
      emailAddress: data.emailAddress.value,
      password: data.password.value,
      firstName: data.firstName,
      lastName: data.lastName,
      createdAt: TimeUtil.toDate(data.createdAt),
      ...(data.updatedAt && { updatedAt: TimeUtil.toDate(data.updatedAt) }),
    };
  }
}
