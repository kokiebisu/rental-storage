import { Payment } from "../../domain/model";

export interface PaymentRepository {
  save(data: Payment): Promise<Payment>;
}
