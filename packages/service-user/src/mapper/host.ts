import { HostRawInterface } from "../entity";

export class HostMapper {
  public static toDTOFromRaw(data: HostRawInterface) {
    return {
      id: data.id,
      firstName: data.first_name,
      lastName: data.last_name,
    };
  }
}
