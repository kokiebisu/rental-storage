import { Payment } from "../../domain/Model";

export interface PaymentRepository {
  save(data: Payment): Promise<Payment>;
}
