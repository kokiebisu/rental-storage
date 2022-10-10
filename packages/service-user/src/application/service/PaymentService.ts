import axios from "axios";
import { PaymentRepositoryImpl } from "../../adapter/out/db/PaymentRepository";
import { Payment } from "../../domain/model";
import { LoggerUtil } from "../../utils";
import { AddPaymentInput, PaymentRepository } from "../port";
import { PaymentService } from "../port";

export class PaymentServiceImpl implements PaymentService {
  private _paymentRepository: PaymentRepository;

  private _logger: LoggerUtil;

  private constructor(paymentRepository: PaymentRepository) {
    this._paymentRepository = paymentRepository;
    this._logger = new LoggerUtil("PaymentServiceImpl");
  }

  public static async create() {
    const paymentRepository = await PaymentRepositoryImpl.create();
    await paymentRepository.setup();
    return new PaymentServiceImpl(paymentRepository);
  }

  public async addPayment(data: AddPaymentInput) {
    this._logger.info(data, "addPayment()");
    // get stripe customer id for the user
    try {
      const response: { data: { providerId: string; providerType: string } } =
        await axios.post(
          `${process.env.SERVICE_API_ENDPOINT}/payments/customer`,
          {
            emailAddress: data.emailAddress,
            firstName: data.firstName,
            lastName: data.lastName,
          }
        );
      console.log("RESPONSE: ", response);
      const payment = new Payment({
        userId: data.userId,
        providerType: response.data.providerType,
        providerId: response.data.providerId,
      });
      this._paymentRepository.save(payment);
    } catch (err) {
      console.error(err);
    }
  }
}
