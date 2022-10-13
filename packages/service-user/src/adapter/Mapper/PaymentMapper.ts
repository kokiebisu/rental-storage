import { Payment } from "../../domain/Model";
import { PaymentInterface } from "../../types";


export class PaymentMapper {
  public static toDTOFromEntity(entity: Payment): PaymentInterface {
    return {
      ...(entity.id && { id: entity.id }),
      providerId: entity.providerId,
      providerType: entity.providerType,
      userId: entity.userId,
    };
  }
}
