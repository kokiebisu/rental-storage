export interface AddPaymentInput {
  userId: number;
  emailAddress: string;
  firstName: string;
  lastName: string;
}

export interface PaymentService {
  addPayment(data: AddPaymentInput): Promise<void>;
}
