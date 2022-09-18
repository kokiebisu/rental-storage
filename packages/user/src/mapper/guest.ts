import { GuestRawInterface } from "../entity";

export class GuestMapper {
  public static toDTOFromRaw(data: GuestRawInterface) {
    return {
      id: data.id,
      firstName: data.first_name,
      lastName: data.last_name,
    };
  }
}
