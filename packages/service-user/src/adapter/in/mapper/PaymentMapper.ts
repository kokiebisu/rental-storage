import { Payment } from "../../../domain/model";
import { PaymentInterface } from "../../../types";

export class PaymentMapper {
  public static toDTOFromEntity(entity: Payment): PaymentInterface {
    return {
      ...(entity.id && { id: entity.id }),
      customerId: entity.customerId,
      uid: entity.uid,
      providerType: entity.providerType,
      userId: entity.userId,
    };
  }
}
