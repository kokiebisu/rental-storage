import { Guest } from "../../../domain/model";
import { GuestInterface, GuestRawInterface } from "../../../types";

export class GuestMapper {
  public static toDTOFromRaw(data: GuestRawInterface) {
    return {
      id: data.id,
      firstName: data.first_name,
      lastName: data.last_name,
      items: data.items,
    };
  }

  public static toDTOFromEntity(data: Guest): GuestInterface {
    return {
      ...(data.id && { id: data.id }),
      firstName: data.firstName,
      lastName: data.lastName,
      items: data.items,
    };
  }
}
