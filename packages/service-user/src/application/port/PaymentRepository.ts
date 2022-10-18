import { Payment } from "../../Domain/Model";

export interface PaymentRepository {
  save(data: Payment): Promise<Payment>;
}
