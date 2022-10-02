import { Guest } from "../../../domain/model";
import { GuestInterface, GuestRawInterface } from "../../../types";
import { TimeUtil } from "../../../utils";

export class GuestMapper {
  public static toDTOFromRaw(data: GuestRawInterface): GuestInterface {
    return {
      id: data.id,
      uid: data.uid,
      emailAddress: data.email_address,
      password: data.password,
      firstName: data.first_name,
      lastName: data.last_name,
      items: data.items,
      createdAt: data.created_at,
      ...(data.updated_at && { updatedAt: data.updated_at }),
    };
  }

  public static toDTOFromEntity(data: Guest): GuestInterface {
    return {
      id: data.id,
      uid: data.uid,
      emailAddress: data.emailAddress.value,
      password: data.password.value,
      firstName: data.firstName,
      lastName: data.lastName,
      items: data.items,
      createdAt: TimeUtil.toDate(data.createdAt),
      ...(data.updatedAt && {
        updatedAt: TimeUtil.toDate(data.updatedAt),
      }),
    };
  }
}
