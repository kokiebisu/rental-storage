import { User } from "../../../domain/model";
import { UserInterface, UserRawInterface } from "../../../types";
import { TimeUtil } from "../../../utils";

export class UserMapper {
  public static toDTOFromRaw(data: UserRawInterface): UserInterface {
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

  public static toDTOFromEntity(data: User): UserInterface {
    return {
      id: data.id,
      uid: data.uid,
      emailAddress: data.emailAddress.value,
      password: data.password,
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
