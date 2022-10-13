import { EmailAddress, Name, NameType, Payment, User } from "../../domain/Model";
import { UserInterface, UserRawInterface } from "../../types";
import { TimeUtil } from "../../utils";


export class UserMapper {
  public static toEntityFromRaw(data: UserRawInterface): User {
    return new User({
      id: data.id,
      uid: data.uid,
      emailAddress: new EmailAddress(data.email_address),
      password: data.password,
      firstName: new Name(NameType.FirstName, data.first_name),
      lastName: new Name(NameType.LastName, data.last_name),
      ...(data.payment_id &&
        data.payment_provider_id &&
        data.payment_provider_type && {
          payment: new Payment({
            id: data.payment_id,
            userId: data.id,
            providerId: data.payment_provider_id,
            providerType: data.payment_provider_type,
          }),
        }),
    });
  }

  public static toDTOFromEntity(data: User): UserInterface {
    return {
      id: data.id,
      uid: data.uid,
      emailAddress: data.emailAddress.value,
      password: data.password,
      firstName: data.firstName.value,
      lastName: data.lastName.value,
      items: data.items,
      ...(data.payment && {
        payment: {
          providerId: data.payment?.providerId,
          providerType: data.payment?.providerType,
        },
      }),
      createdAt: TimeUtil.toDate(data.createdAt),
      ...(data.updatedAt && {
        updatedAt: TimeUtil.toDate(data.updatedAt),
      }),
    };
  }
}
