import { Payment } from "../../Domain/Model";
import { PaymentInterface } from "../../Types";

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
